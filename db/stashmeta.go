package db

import (
	"fmt"

	"github.com/boltdb/bolt"
)

const stashBucket string = "stashes"

// getStashMetaBucket returns the bucket corresponding
// to a specific league for holding stash metadata
//
// Will either panic or return a valid bucket.
func getStashMetaBucket(league LeagueHeapID, tx *bolt.Tx) *bolt.Bucket {
	// Grab league bucket
	leagueBucket := getLeagueBucket(league, tx)

	// Grab the meta
	//
	// This can never fail, its a guarantee that the stashBucket was registered
	// and will always appear on a valid leagueBucket
	metaBucket := leagueBucket.Bucket([]byte(stashBucket))
	if metaBucket == nil {
		panic(fmt.Sprintf("%s bucket not found when expected", stashBucket))
	}

	return metaBucket
}

// AddStashes adds tbe given items to their correct paths in the database
//
// Provided stashes CAN differ in their league.
func AddStashes(stashes []Stash, items [][]Item, db *bolt.DB) (int, error) {

	// Silently exit when no items stashes to add
	if len(stashes) < 1 {
		return 0, nil
	}
	if len(stashes) != len(items) {
		return 0, fmt.Errorf("each stash must have matching items, got %d!=%d",
			len(stashes), len(items))
	}

	// Keep track of the number of items we overwrite in this process
	updated := 0

	return updated, db.Update(func(tx *bolt.Tx) error {

		// Add all of the stash metadata to the stashMeta
		for i, stash := range stashes {

			// Serialize the item
			serial, err := stash.MarshalMsg(nil)
			if err != nil {
				return fmt.Errorf("failed to Marshal Stash, err=%s", err)
			}

			meta := getStashMetaBucket(stash.League, tx)

			// Check for a pre-existing stash
			val := meta.Get(stash.ID[:])
			if val != nil {
				// TODO: handle diffing a stash
				fmt.Println("I should've diffed and updated a stash... but I didn't...")
				updated++
			}

			meta.Put(stash.ID[:], serial)

			// Add the items for this stash
			if _, err := addItems(items[i], stash.When, tx); err != nil {
				return fmt.Errorf("failed to add items for stash id=%s, err=%s",
					stash.ID, err)
			}
		}
		return nil
	})

}
