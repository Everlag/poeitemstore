package db

import (
	"fmt"

	"github.com/boltdb/bolt"
)

const itemStoreBucket string = "itemStore"

// getLeagueBucket returns the bucket corresponding
// to a specific league
//
// Will either panic or return a valid bucket.
func getLeagueBucket(league StringHeapID, tx *bolt.Tx) *bolt.Bucket {
	// Grab root Bucket
	rootBucket := tx.Bucket([]byte(itemStoreBucket))
	if rootBucket == nil {
		panic(fmt.Sprintf("root item store bucket not found %s, ", itemStoreBucket))
	}

	leagueBytes := StringHeapIDToBytes(league)
	leagueBucket := rootBucket.Bucket(leagueBytes)
	if leagueBucket == nil {
		var err error
		leagueBucket, err = rootBucket.CreateBucket(leagueBytes)
		if err != nil {
			panic(fmt.Sprintf("cannot create league bucket, err=%s", err))
		}
	}
	return leagueBucket
}

// AddItems adds tbe given items to their correct paths in the database
//
// Provided items CAN differ in their league.
func AddItems(items []Item, db *bolt.DB) error {

	// Silently exit when no items present to add
	if len(items) < 1 {
		return nil
	}

	// Allocate a buffer to place our serialized items
	//
	// This allows us to avoid a whole ton of allocations
	//
	// We know, at least, items[0] exists due to len check at start
	prealloc := make([]byte, items[0].Msgsize())

	return db.Update(func(tx *bolt.Tx) error {

		for _, item := range items {
			// Reset slice to zero length but keep capacity to avoid allocations
			prealloc = prealloc[:1]

			// Serialize the item
			serial, err := item.MarshalMsg(prealloc)
			if err != nil {
				return fmt.Errorf("failed to Marshal Item, err=%s", err)
			}

			league := getLeagueBucket(item.League, tx)

			league.Put(item.ID[:], serial)
		}

		return nil
	})

}
