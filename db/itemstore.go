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
func addItems(items []Item, tx *bolt.Tx) (int, error) {

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

	// Add all of the items to the itemStore
	for _, item := range items {

		// Serialize the item
		serial, err := item.MarshalMsg(nil)
		if err != nil {
			return 0, fmt.Errorf("failed to Marshal Item, err=%s", err)
		}

		league := getLeagueItemBucket(item.League, tx)

		val := league.Get(item.ID[:])
		if val != nil {
			overwritten++
		}

		league.Put(item.ID[:], serial)
	}

	// Index each of the items
	_, err := IndexItems(items, tx)
	if err != nil {
		return 0, fmt.Errorf("failed to add indices, err=%s", err)
	}

	return overwritten, nil

}

// AddItems adds tbe given items to their correct paths in the database
//
// Provided items CAN differ in their league.
//
// This is simply a wrapper for addItems that also establishes the transaction
// as well as include the timestamp
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
		var err error
		overwritten, err = addItems(items, tx)
		return err
	})

}

// RemoveItems removes the provided IDs from the database
//
// All provided IDs must be under the same League. Timestamps are provided
// on a per-item basis to allow index removal
//
// For higher performance, callers can sort the IDs to allow for more
// sequential access behaviors when handling the heap
func RemoveItems(ids []ID, league LeagueHeapID, db *bolt.DB) error {

	return db.Update(func(tx *bolt.Tx) error {
		// Get the league bucket
		leagueBucket := getLeagueItemBucket(league, tx)

		// Keep track of details for removed items so they can have their
		// indices removed
		items := make([]Item, len(ids))
		for i, id := range ids {
			// We need to fetch the item to have sufficient information
			// to remove all of its index entries
			itemBytes := leagueBucket.Get(id[:])
			if itemBytes == nil {
				return fmt.Errorf("item not found, cannot remove")
			}
			// Delete the item from the heap
			leagueBucket.Delete(id[:])

			var item Item
			if _, err := item.UnmarshalMsg(itemBytes); err != nil {
				return fmt.Errorf("failed to Unmarshal Item from heap, err=%s", err)
			}
			items[i] = item
		}

		// Remove associate index entries
		if err := DeindexItems(items, tx); err != nil {
			return fmt.Errorf("failed remove item indices, err=%s", err)
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
