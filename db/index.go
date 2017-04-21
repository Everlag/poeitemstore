package db

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
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
	for _, key := range keys {
		keyBytes := key.ToBytes()
		prevBucket := currentBucket.Bucket(keyBytes)
		if prevBucket == nil {
			// Create the bucket
			var err error
			prevBucket, err = currentBucket.CreateBucket(keyBytes)
			if err != nil {
				return nil,
					errors.Wrapf(err,
						"failed to add index intermediary bucket bucket, bucket=%s, chain=%v",
						key, keys)
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
			return nil, errors.Errorf("invalid bucket, key=%d, chain=%v", key, keys)
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

	// Pre-allocate index key so the entire key can be
	// encoded with a single allocation.
	modsLength := 2 * len(mod.Values)
	indexKey := make([]byte, ModIndexKeySuffixLength+modsLength)

	// Generate the suffix
	sequenceBytes := i16tob(updateSequence)
	suffix := (indexKey[modsLength:])[:0] // Deal with pre-allocated space
	suffix = append(suffix, now[:]...)
	suffix = append(suffix, sequenceBytes...)

	if len(suffix) != ModIndexKeySuffixLength {
		panic(fmt.Sprintf("unexpected suffix length, got %d, expected %d",
			len(suffix), ModIndexKeySuffixLength))
	}

	// Fill in the index from the front
	//
	// TODO: avoid appends, pre-size the backing slice to accomodate the
	// contents including the header
	index := indexKey[:0] // Deal with pre-allocated space
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
		return nil, errors.New("invalid index key passed, less than length of suffix")
	}

	// Ensure we are divisible by 2 following the removal of the suffix
	if (len(key)-ModIndexKeySuffixLength)%2 != 0 {
		return nil, errors.New("invalid index key passed, values malformed")
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
func IndexItems(items []Item, tx *bolt.Tx) (int, error) {

	// Sanity check passed in transaction, better to do this than panic.
	if !tx.Writable() {
		return 0, errors.New("cannot IndexItems on readonly transaction")
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
				return 0, errors.New("failed to get item mod bucket")
			}

			modKey := encodeModIndexKey(mod, item.When, item.UpdateSequence)

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

// DeindexItems removes tbe given items from their correct indices
//
// If an index entry cannot be removed, we return an error. This ensures
// all existing index entries must be alive
func DeindexItems(items []Item, tx *bolt.Tx) error {

	// Sanity check passed in transaction, better to do this than panic.
	if !tx.Writable() {
		return errors.New("cannot IndexItems on readonly transaction")
	}

	// Silently exit when no items present to add
	if len(items) < 1 {
		return nil
	}

	for _, item := range items {

		for _, mod := range item.Mods {
			// Grab the bucket we can actually insert things into
			itemModBucket, err := getItemModIndexBucket(item.RootType, item.RootFlavor,
				mod.Mod, item.League, tx)
			if err != nil {
				return errors.New("failed to get item mod bucket")
			}

			modKey := encodeModIndexKey(mod, item.When, item.UpdateSequence)

			// We need to make a copy of the item ID or bolt
			// will get a buffer reused for all items.
			//
			// Without this, all index entries will point to the last
			// item added.
			idCopy := make([]byte, IDSize)
			copy(idCopy, item.ID[:])

			itemModBucket.Delete(modKey)
		}
	}

	return nil

}

// IndexEntryCount returns the number of index entries across all leagues
func IndexEntryCount(db *bolt.DB) (int, error) {
	var count int

	leagueStrings, err := ListLeagues(db)
	if err != nil {
		return 0, err
	}
	leagueIDs, err := GetLeagues(leagueStrings, db)
	if err != nil {
		return 0, err
	}

	return count, db.View(func(tx *bolt.Tx) error {

		for _, id := range leagueIDs {
			b := getLeagueIndexBucket(id, tx)
			if b == nil {
				return errors.Errorf("%s bucket not found", itemStoreBucket)
			}
			stats := b.Stats()
			count += stats.KeyN
		}

		return nil
	})
}
