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

// getNextItemID returns the next identifier an item in the provided
// league item heap bucke.
func getNextItemID(league *bolt.Bucket) (ID, error) {

	// If it doesn't, we need a sequence number
	seq, err := league.NextSequence()
	if err != nil {
		return ID{}, fmt.Errorf("failed to get NextSequence in %s",
			itemStoreBucket)
	}
	return IDFromSequence(seq), nil
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

	// Acquire now
	now := NewTimestamp()

	return overwritten, db.Update(func(tx *bolt.Tx) error {

		// Add all of the items to the itemStore
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

		// Index each of the items
		_, err := IndexItems(items, now, tx)
		if err != nil {
			return fmt.Errorf("failed to add indices, err=%s", err)
		}

		return nil
	})

}

// GetItemByID returns the item represented by provided ID
// in a specific league
func GetItemByID(id ID, league LeagueHeapID, tx *bolt.Tx) (Item, error) {
	var item Item

	// Get the league bucket
	leagueBucket := getLeagueItemBucket(league, tx)

	// Grab the stored item
	itemBytes := leagueBucket.Get(id[:])
	if itemBytes == nil {
		return item, fmt.Errorf("item not found")
	}

	// Unmarshal the item
	_, err := item.UnmarshalMsg(itemBytes)
	return item, err

}

// GetItemByIDGlobal attempts to resolve an item ID across
// every available league.
func GetItemByIDGlobal(id ID, db *bolt.DB) (Item, error) {
	var item Item

	leagueStrings, err := ListLeagues(db)
	if err != nil {
		return item, err
	}
	leagueIDs, err := GetLeagues(leagueStrings, db)
	if err != nil {
		return item, err
	}

	fmt.Printf("got league ids %v\n", leagueIDs)

	return item, db.View(func(tx *bolt.Tx) error {

		var err error
		for _, league := range leagueIDs {
			item, err = GetItemByID(id, league, tx)
			if err == nil {
				return nil
			}
		}

		return err
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
