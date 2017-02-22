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
func getLeagueBucket(league LeagueHeapID, tx *bolt.Tx) *bolt.Bucket {
	// Grab root Bucket
	rootBucket := tx.Bucket([]byte(itemStoreBucket))
	if rootBucket == nil {
		panic(fmt.Sprintf("root item store bucket not found %s, ", itemStoreBucket))
	}

	leagueBytes := LeagueHeapIDToBytes(league)
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

	// NOTE: we cannot preallocate the buffer for marshalling
	// each item as the buffer must remain constant post-Put
	// for the duration of the transaction.
	//
	// So yeah, that was fun to figure out. RTFM helped :|

	return db.Update(func(tx *bolt.Tx) error {

		for _, item := range items {

			// Serialize the item
			serial, err := item.MarshalMsg(nil)
			if err != nil {
				return fmt.Errorf("failed to Marshal Item, err=%s", err)
			}

			league := getLeagueBucket(item.League, tx)

			val := league.Get(item.ID[:])
			if val != nil {
				fmt.Println("overwriting val...")
			}

			league.Put(item.ID[:], serial)
		}

		return nil
	})

}

// // GetItems returns a compact Item for each ID provided
// func GetItems(ids []ID, db *bolt.DB) []Item  {

// }

// ItemStoreCount returns the number of items across all leagues
func ItemStoreCount(db *bolt.DB) (int, error) {
	var count int

	return count, db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(itemStoreBucket))
		if b == nil {
			return fmt.Errorf("%s bucket not found", itemStoreBucket)
		}

		stats := b.Stats()
		count = stats.KeyN

		return nil
	})
}
