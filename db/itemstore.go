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
func getLeagueItemBucket(league LeagueHeapID, tx *bolt.Tx) *bolt.Bucket {
	// Grab league bucket
	leagueBucket := getLeagueBucket(league, tx)

	// Grab the itemStore
	//
	// This can never fail, its a guarantee that the itemStoreBucket was registered
	// and will always appear on a valid leagueBucket
	itemStore := leagueBucket.Bucket([]byte(itemStoreBucket))
	if itemStore == nil {
		panic(fmt.Sprintf("%s bucket not found when expected", itemStoreBucket))
	}

	return itemStore
}

// AddItems adds tbe given items to their correct paths in the database
//
// Provided items CAN differ in their league.
func AddItems(items []Item, db *bolt.DB) (int, error) {

	// Silently exit when no items present to add
	if len(items) < 1 {
		return 0, nil
	}

	// NOTE: we cannot preallocate the buffer for marshalling
	// each item as the buffer must remain constant post-Put
	// for the duration of the transaction.
	//
	// So yeah, that was fun to figure out. RTFM helped :|

	// Keep track of the number of items we overwrite in this process
	overwritten := 0

	return overwritten, db.Update(func(tx *bolt.Tx) error {

		for _, item := range items {

			// Serialize the item
			serial, err := item.MarshalMsg(nil)
			if err != nil {
				return fmt.Errorf("failed to Marshal Item, err=%s", err)
			}

			league := getLeagueItemBucket(item.League, tx)

			val := league.Get(item.ID[:])
			if val != nil {
				overwritten++
			}

			league.Put(item.ID[:], serial)
		}

		return nil
	})

}

// ItemStoreCount returns the number of items across all leagues
func ItemStoreCount(db *bolt.DB) (int, error) {
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
			b := getLeagueItemBucket(id, tx)
			if b == nil {
				return fmt.Errorf("%s bucket not found", itemStoreBucket)
			}
			stats := b.Stats()
			count += stats.KeyN
		}

		return nil
	})
}
