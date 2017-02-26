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
