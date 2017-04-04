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

// stashDiffUpdate handles updating a stash's stored items
// by diffing it then adding and removing items which are new and
// expired, respectively.
//
// This DOES NOT actually update the stashmeta bucket entry for the stash
func stashDiffUpdate(oldSerial []byte, newStash Stash, newItems []Item,
	tx *bolt.Tx) error {

	var old Stash
	if _, err := old.UnmarshalMsg(oldSerial); err != nil {
		return fmt.Errorf("failed to unmarshal oldSerial")
	}

	// Determine which items get added and which get removed
	add, remove := newStash.Diff(old)

	// Translate the ids to remove
	removeIDs, err := GetGGGIDTranslations(remove, newStash.League, tx)
	if err != nil {
		return fmt.Errorf("failed to translate GGGIDs to local IDs, err=%s",
			err)
	}
	// And get rid of them
	err = removeItems(removeIDs, newStash.League, tx)
	if err != nil {
		return fmt.Errorf("failed to removeItems, err=%s", err)
	}

	// Filter out the Items to add from the items contained in this update
	toAddFilter := make(map[GGGID]struct{})
	toAdd := make([]Item, 0)
	for _, id := range add {
		toAddFilter[id] = struct{}{}
	}
	for _, item := range newItems {
		if _, ok := toAddFilter[item.GGGID]; ok {
			toAdd = append(toAdd, item)
		}
	}

	// Add the items
	if _, err := addItems(toAdd, tx); err != nil {
		return fmt.Errorf("failed to addItems, err=%s", err)
	}

	return nil

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

			// Serialize the stash
			serial, err := stash.MarshalMsg(nil)
			if err != nil {
				return fmt.Errorf("failed to Marshal Stash, err=%s", err)
			}

			meta := getStashMetaBucket(stash.League, tx)

			// Check for a pre-existing stash
			oldSerial := meta.Get(stash.ID[:])
			if oldSerial == nil {
				// Handle trivial case of just needing to add the entire stash
				// Add the items for this stash
				if _, err := addItems(items[i], tx); err != nil {
					return fmt.Errorf("failed to add items for stash id=%s, err=%s",
						stash.ID, err)
				}
			} else {
				// Handle trivial case of just needing to add the entire stash
				meta.Put(stash.ID[:], serial)
				// Take care of diffing the stash
				stashDiffUpdate(oldSerial, stash, items[i], tx)
				updated++
			}

			// Then update the metadata
			meta.Put(stash.ID[:], serial)

		}
		return nil
	})

}
