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
func getItemModIndexBucket(mod ItemMod, item Item, tx *bolt.Tx) (*bolt.Bucket, error) {
	// Keys towards the bucket we want to return, they may or may not exist
	keys := []StringHeapID{item.RootType, item.RootFlavor, mod.Mod}

	// Start at the index bucket
	currentBucket := getLeagueIndexBucket(item.League, tx)

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

// encodeModIndexKey generates a mod key based off of the provided data
//
// The mod index key is generated as [mod.Values..., now, updateSequence]
func encodeModIndexKey(mod ItemMod, now Timestamp, updateSequence uint16) []byte {
	// Generate the suffix
	suffix := make([]byte, 0)
	suffix = append(suffix, now[:]...)
	suffix = append(suffix, i16tob(updateSequence)...)

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

// TODO
// //
// // Okay, so we have the itemIndexBucket, now we need to have a helper function
// // that will find, and potentially create, the appropriate modifier bucket.
// //
// // From there, we'll take that modifer bucket, serialize the modifier to a key
// // of something like key=[values, currentTimeTruncatedToDays(???)], value=itemID

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
			itemModBucket, err := getItemModIndexBucket(mod, item, tx)
			if err != nil {
				return 0, fmt.Errorf("failed to get item mod bucket")
			}

			modKey := encodeModIndexKey(mod, now, item.UpdateSequence)

			itemModBucket.Put(modKey, item.ID[:])
			added++
		}
	}

	return added, nil

}
