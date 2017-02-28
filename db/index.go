package db

import (
	"fmt"

	"github.com/boltdb/bolt"
)

const indiceBucket string = "indices"

// getLeagueIndexBucketRO returns the bucket corresponding
// to a specific league's index. This will never write
// and can be used safely with a readonly transaction.
//
// Will either panic or return a valid bucket.
func getLeagueIndexBucket(league LeagueHeapID, tx *bolt.Tx) *bolt.Bucket {
	// Grab league bucket
	leagueBucket := getLeagueBucket(league, tx)

	// This can never fail, its a guarantee that the itemStoreBucket was registered
	// and will always appear on a valid leagueBucket
	indices := leagueBucket.Bucket([]byte(indiceBucket))
	if indices == nil {
		panic(fmt.Sprintf("%s bucket not found when expected", itemStoreBucket))
	}

	return indices
}

// getItemModIndexBucket returns a bucket which a given mod can be put
// when considering the item containing it.
//
// This WILL write if a bucket is not found. Hence, readonly tx unsafe.
func getItemModIndexBucket(rootType, rootFlavor, mod StringHeapID,
	league LeagueHeapID, tx *bolt.Tx) (*bolt.Bucket, error) {
	// Keys towards the bucket we want to return, they may or may not exist
	keys := []StringHeapID{rootType, rootFlavor, mod}

	// Start at the index bucket
	currentBucket := getLeagueIndexBucket(league, tx)

	// Create all of the intervening keys
	for i, key := range keys {
		keyBytes := key.ToBytes()
		prevBucket := currentBucket.Bucket(keyBytes)
		if prevBucket == nil {
			// Create the bucket
			var err error
			if i == 0 {
				// fmt.Println("creating non-existent bucket!", keyBytes)
			}
			prevBucket, err = currentBucket.CreateBucket(keyBytes)
			if err != nil {
				return nil,
					fmt.Errorf("failed to add index intermediary bucket bucket, bucket=%s, chain=%v, err=%s",
						key, keys, err)
			}
		}
		currentBucket = prevBucket
	}

	// If we made it through, our currentBucket should be the one we want
	return currentBucket, nil
}

// getItemModIndexBucketRO returns a bucket which a given mod can be put
// when considering the item containing it.
//
// This WILL NOT write if a bucket is not found. Hence, readonly tx unsafe.
func getItemModIndexBucketRO(rootType, rootFlavor, mod StringHeapID,
	league LeagueHeapID, tx *bolt.Tx) (*bolt.Bucket, error) {
	// Keys towards the bucket we want to return, they may or may not exist
	keys := []StringHeapID{rootType, rootFlavor, mod}

	// Start at the index bucket
	currentBucket := getLeagueIndexBucket(league, tx)

	// Traverse all intervening buckets
	for _, key := range keys {
		keyBytes := key.ToBytes()
		prevBucket := currentBucket.Bucket(keyBytes)
		if prevBucket == nil {
			return nil, fmt.Errorf("invalid bucket, key=%d, chain=%v", key, keys)
		}
		currentBucket = prevBucket
	}

	// If we made it through, our currentBucket should be the one we want
	return currentBucket, nil
}

// ModIndexKeySuffixLength allows us to fetch variable numbers
// of pre-pended values given their length.
const ModIndexKeySuffixLength = TimestampSize + 2

// encodeModIndexKey generates a mod key based off of the provided data
//
// The mod index key is generated as [mod.Values..., now, updateSequence]
func encodeModIndexKey(mod ItemMod, now Timestamp, updateSequence uint16) []byte {
	// Generate the suffix
	suffix := make([]byte, 0)
	suffix = append(suffix, now[:]...)
	suffix = append(suffix, i16tob(updateSequence)...)

	if len(suffix) != ModIndexKeySuffixLength {
		panic(fmt.Sprintf("unexpected suffix length, got %d, expected %d",
			len(suffix), ModIndexKeySuffixLength))
	}

	// Fill in the index from the front
	//
	// TODO: avoid appends, pre-size the backing slice to accomodate the
	// contents including the header
	index := make([]byte, 0)
	for _, value := range mod.Values {
		index = append(index, i16tob(value)...)
	}

	// And return the index with its suffix
	return append(index, suffix...)
}

// decodeModIndexKey decodes a provided mod index key
//
// This returns the values encoded in the key.
//
// This is possible as the suffix is a fixed length and format while
// the values of the modifer are simple appended
func decodeModIndexKey(key []byte) ([]uint16, error) {

	// Basic sanity check
	if len(key) < ModIndexKeySuffixLength {
		return nil, fmt.Errorf("invalid index key passed, less than length of suffix")
	}

	// Ensure we are divisible by 2 following the removal of the suffix
	if (len(key)-ModIndexKeySuffixLength)%2 != 0 {
		return nil, fmt.Errorf("invalid index key passed, values malformed")
	}

	valueBytes := key[:len(key)-ModIndexKeySuffixLength]
	values := make([]uint16, len(valueBytes)/2)
	for index := 0; index*2 < len(valueBytes); index++ {
		values[index] = btoi16(valueBytes[index*2:])
	}

	return values, nil

}

// IndexItems adds tbe given items to their correct indices
// for efficient lookup. Returns number of index entries added.
//
// Provided items CAN differ in their league.
func IndexItems(items []Item, now Timestamp, tx *bolt.Tx) (int, error) {

	// Sanity check passed in transaction, better to do this than panic.
	if !tx.Writable() {
		return 0, fmt.Errorf("cannot IndexItems on readonly transaction")
	}

	// Silently exit when no items present to add
	if len(items) < 1 {
		return 0, nil
	}

	var added int

	for _, item := range items {

		for _, mod := range item.Mods {
			// Grab the bucket we can actually insert things into

			itemModBucket, err := getItemModIndexBucket(item.RootType, item.RootFlavor,
				mod.Mod, item.League, tx)
			if err != nil {
				return 0, fmt.Errorf("failed to get item mod bucket")
			}

			modKey := encodeModIndexKey(mod, now, item.UpdateSequence)

			// We need to make a copy of the item ID or bolt
			// will get a buffer reused for all items.
			//
			// Without this, all index entries will point to the last
			// item added.
			idCopy := make([]byte, IDSize)
			copy(idCopy, item.ID[:])

			itemModBucket.Put(modKey, idCopy)
			added++
		}
	}

	return added, nil

}

// LookupItems returns up to n item IDs where those items
// are of the given type and flavor while also containing
// the provided mod with a given minimum value
//
// Right now minModValue checks only against the first value
// if a mod contains multiple values, this can be an array
// in the future. We could even validate this by checking
// the mod StringHeapID to ensure it has the right number of
// placeholder signs.
// TODO: review using array for minModValue
func LookupItems(rootType, rootFlavor, mod StringHeapID,
	league LeagueHeapID,
	minModValue uint16,
	n int, db *bolt.DB) ([]ID, error) {

	// ids presized...
	ids := make([]ID, n)
	lastIDFound := 0

	err := db.View(func(tx *bolt.Tx) error {

		itemModBucket, err := getItemModIndexBucketRO(rootType, rootFlavor, mod, league, tx)
		if err != nil {
			return fmt.Errorf("faield to get item mod index bucket, err=%s", err)
		}

		// Grab a cursor
		c := itemModBucket.Cursor()

		// Iterate over items in reverse sorted key order. This starts
		// from the last key/value pair and updates the k/v variables to
		// the previous key/value on each iteration.
		//
		// The loop finishes at the beginning of the cursor when a nil key
		// is returned.
		for k, v := c.Last(); k != nil; k, v = c.Prev() {
			// Ignore nested buckets
			if k == nil {
				continue
			}

			// Decode the key
			values, err := decodeModIndexKey(k)
			if err != nil {
				return fmt.Errorf("failed to decode item mod index key, key=%v, err=%s",
					k, err)
			}
			if len(values) == 0 {
				return fmt.Errorf("decoded item mod index key to no values, key=%v", k)
			}

			// Ensure the mod is the correct value
			if values[0] >= minModValue {
				if len(v) != IDSize {
					panic(fmt.Sprintf("malformed id value in index, incorrect length; id=%v", v))
				}
				var id ID
				copy(id[:], v)
				ids[lastIDFound] = id
				lastIDFound++
				// Check if we're done
				if lastIDFound >= n {
					return nil
				}
			}
		}

		return nil
	})

	// Truncate ids to however many we actually found
	ids = ids[:lastIDFound]

	return ids, err
}

// LookupItemsMultiModStrideLength determines how many items
// is included in a stride of LookupItemsMultiMod.
//
// Longer strides mean fewer intersections but more potentially useless
// item mods checked.
const LookupItemsMultiModStrideLength = 32

// lookupMultiModStride performs a single stride using the given cursors
// while filling the given sets if the cursors have the currect value
func lookupMultiModStride(minModValues []uint16,
	sets []map[ID]struct{},
	cursors []*bolt.Cursor, validCursors *int) error {

	// TODO: invert the for loop, we should be iterating along the cursors
	// LookupItemsMultiModStrideLength times rather than this inefficient shit
	for index := 0; index < LookupItemsMultiModStrideLength; index++ {
		for i, c := range cursors {
			// Handle nil cursor indicating that mod
			// has no more legitimate values
			if c == nil {
				continue
			}

			// Grab a pair
			k, v := c.Prev()
			// Ignore nested buckets
			if k == nil {
				continue
			}
			// Grab the value
			values, err := decodeModIndexKey(k)
			if err != nil {
				return fmt.Errorf("failed to decode mod index key, key=%v, err=%s",
					k, err)
			}
			if len(values) == 0 {
				return fmt.Errorf("decoded item mod index key to no values, key=%v", k)
			}

			// Ensure the mod is the correct value
			if values[0] >= minModValues[i] {
				if len(v) != IDSize {
					panic(fmt.Sprintf("malformed id value in index, incorrect length; id=%v", v))
				}
				// NOTE: the copy here is actually completely required
				// due to the fact that boltdb makes no guarantee regarding what
				// keys and value slices contain when outside a transaction.
				var id ID
				copy(id[:], v)
				sets[i][id] = struct{}{}
			} else {
				// Remove from cursors we're interested in
				cursors[i] = nil
				*validCursors--
			}
		}
	}
	return nil
}

// intersectItemIDMaps returns how many items in the given sets appear
// across all them. The found items have their IDs put in result
// up to its length.
//
// Pass a nil result to obtain just the count
func intersectItemIDSets(sets []map[ID]struct{}, result []ID) int {

	// Keep track of our maximum matches
	//
	// When nil result, we account for that in logic
	n := len(result)

	// And how many we have found so far
	foundIDs := 0

	// Intersect the sets by taking one of them
	// and seeing how many of its items appear in others
	firstSet := sets[0]
	for id := range firstSet {
		// sharedCount always starts at one because it
		//  is always shared with the firstSet
		sharedCount := 1
		for _, other := range sets[1:] {
			_, ok := other[id]
			if ok {
				sharedCount++
			}
		}
		if sharedCount == len(sets) {
			foundIDs++
			// Add the item if we need to
			if result != nil {
				result[foundIDs-1] = id
			}
			// Exit early if we reach capacity
			if result != nil && foundIDs >= n {
				return foundIDs
			}
		}
	}

	return foundIDs
}

// LookupItemsMultiMod returns up to n item IDs where those items
// are of the given type and flavor while also containing
// the provided mods each with their minimum values
func LookupItemsMultiMod(rootType, rootFlavor StringHeapID,
	mods []StringHeapID, minModValues []uint16,
	league LeagueHeapID,
	n int, db *bolt.DB) ([]ID, error) {

	// ids presized...
	var ids []ID

	// The basic strategy here
	// is to find the buckets we need
	// setup sets to handle placing items in
	// iterate over STRIDE_LENGTH items and add them to the sets
	// check if the last stride has found n intersecting items across all sets
	// If yes, done
	// If not, remove any mods whose keys are below their min values and do another stride

	err := db.View(func(tx *bolt.Tx) error {

		// Make a place to keep our cursors
		//
		// NOTE: a cursor can be nil to indicate it should not be queried
		cursors := make([]*bolt.Cursor, len(mods))

		// Keep track of how many cursors are valid,
		// this will let us know when we've exhausted our data
		validCursors := len(cursors)

		// Collect our buckets for each mod and establish cursors
		for i, mod := range mods {
			itemModBucket, err := getItemModIndexBucketRO(rootType, rootFlavor, mod, league, tx)
			if err != nil {
				return fmt.Errorf("faield to get item mod index bucket, mod=%d err=%s",
					mod, err)
			}
			cursors[i] = itemModBucket.Cursor()
		}

		// Create our item sets
		sets := make([]map[ID]struct{}, len(mods))
		for i := range sets {
			sets[i] = make(map[ID]struct{})
		}

		// Set all of our cursors to be at their ends
		for i, c := range cursors {
			// Set to last
			k, v := c.Last()
			// Ignore nested buckets
			if k == nil {
				continue
			}
			// Grab the value
			values, err := decodeModIndexKey(k)
			if err != nil {
				return fmt.Errorf("failed to decode mod index key, err=%s", err)
			}
			if len(values) == 0 {
				return fmt.Errorf("decoded item mod index key to no values, key=%v", k)
			}

			// Ensure the mod is the correct value
			if values[0] >= minModValues[i] {
				if len(v) != IDSize {
					panic(fmt.Sprintf("malformed id value in index, incorrect length; id=%v", v))
				}
				// NOTE: the copy here is actually completely required
				// due to the fact that boltdb makes no guarantee regarding what
				// keys and value slices contain when outside a transaction.
				var id ID
				copy(id[:], v)
				sets[i][id] = struct{}{}
			} else {
				// Remove from cursors we're interested in
				cursors[i] = nil
				validCursors--
			}
		}

		// Perform our strides to search
		var foundIDs int
		for foundIDs < n && validCursors > 0 {
			// Iterate for a stride
			err := lookupMultiModStride(minModValues, sets, cursors, &validCursors)
			if err != nil {
				return fmt.Errorf("failed a stride, err=%s", err)
			}

			foundIDs = intersectItemIDSets(sets, nil)
		}

		// Cap result to desired length as required
		if foundIDs > n {
			foundIDs = n
		}
		// Perform one more intersection to find our return value
		ids = make([]ID, foundIDs)
		intersectItemIDSets(sets, ids)

		return nil
	})

	return ids, err
}
