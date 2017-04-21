package db

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

const stashBucket string = "stashes"

// StashUpdateStats represents the actual
// work done in an operation that is applied to Stashes.
//
// All values are expected to be >= 0
type StashUpdateStats struct {
	Added   int // Number of stashes added
	Updated int // Number of stashes updated
	Intact  int // Number of stashes without any changes made
	// Items give item-wise stats
	Items ItemUpdateStats
}

// ItemUpdateStats is a subfield of StashUpdateStats broken
// out for easier instantiation.
type ItemUpdateStats struct {
	Added   int // Number of items added
	Removed int // Number of items removed
	Kept    int // Number of items kept
}

// Compare considers the receiver as the expected StashUpdateStats
// while the other is tested to see if it matches. Any difference
// between expected and other is reported as an error.
func (s *StashUpdateStats) Compare(other *StashUpdateStats) error {
	if s.Added != other.Added {
		return errors.Errorf("mismatched Added, expected %d, got %d",
			s.Added, other.Added)
	}
	if s.Updated != other.Updated {
		return errors.Errorf("mismatched Updated, expected %d, got %d",
			s.Updated, other.Updated)
	}
	if s.Intact != other.Intact {
		return errors.Errorf("mismatched Intact, expected %d, got %d",
			s.Intact, other.Intact)
	}
	if s.Items.Added != other.Items.Added {
		return errors.Errorf("mismatched Items.Added, expected %d, got %d",
			s.Items.Added, other.Items.Added)
	}
	if s.Items.Removed != other.Items.Removed {
		return errors.Errorf("mismatched Items.Removed, expected %d, got %d",
			s.Items.Removed, other.Items.Removed)
	}
	if s.Items.Kept != other.Items.Kept {
		return errors.Errorf("mismatched Items.Kept, expected %d, got %d",
			s.Items.Kept, other.Items.Kept)
	}
	return nil
}

func (s StashUpdateStats) String() string {
	return fmt.Sprintf(`stashes: %d added | %d updated | %d intact
  items: %d added | %d removed | %d kept`,
		s.Added, s.Updated, s.Intact,
		s.Items.Added, s.Items.Removed, s.Items.Kept)
}

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
	stats *StashUpdateStats, tx *bolt.Tx) error {

	var old Stash
	if _, err := old.UnmarshalMsg(oldSerial); err != nil {
		return errors.New("failed to unmarshal oldSerial")
	}

	// Determine which items get added and which get removed
	add, remove := newStash.Diff(old)

	stats.Items.Added += len(add)
	stats.Items.Removed += len(remove)
	stats.Items.Kept += len(newItems) - (len(add) + len(remove))
	// No additions or removals means these are all in the database
	// and currently valid. Hence, we can skip the remainder of our work.
	if len(add)+len(remove) == 0 {
		stats.Intact++
		return nil
	}
	stats.Updated++

	// Translate the ids to remove
	removeIDs, err := GetGGGIDTranslations(remove, newStash.League, tx)
	if err != nil {
		return errors.Wrap(err, "failed to translate GGGIDs to local IDs")
	}
	// And get rid of those to remove
	// TODO: look into sorting these before removal for performance benefits
	err = removeItems(removeIDs, newStash.League, tx)
	if err != nil {
		return errors.Wrap(err, "failed to removeItems")
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
		return errors.Wrap(err, "failed to addItems")
	}

	return nil

}

// AddStashes adds tbe given items to their correct paths in the database
//
// Provided stashes CAN differ in their league.
func AddStashes(stashes []Stash, items [][]Item,
	db *bolt.DB) (*StashUpdateStats, error) {

	// Silently exit when no items stashes to add
	if len(stashes) < 1 {
		return nil, nil
	}
	if len(stashes) != len(items) {
		return nil, errors.Errorf("each stash must have matching items, got %d!=%d",
			len(stashes), len(items))
	}

	stats := StashUpdateStats{}

	return &stats, db.Update(func(tx *bolt.Tx) error {

		// Add all of the stash metadata to the stashMeta
		for i, stash := range stashes {

			// Serialize the stash
			serial, err := stash.MarshalMsg(nil)
			if err != nil {
				return errors.Wrap(err, "failed to Marshal Stash")
			}

			meta := getStashMetaBucket(stash.League, tx)

			// Check for a pre-existing stash
			oldSerial := meta.Get(stash.ID[:])
			if oldSerial == nil {
				// Handle trivial case of just needing to add the entire stash
				// Add the items for this stash
				if _, err := addItems(items[i], tx); err != nil {
					return errors.Wrapf(err, "failed to add items for stash id=%s",
						stash.ID)
				}
				stats.Added++
				stats.Items.Added += len(items[i])
			} else {
				// Handle trivial case of just needing to add the entire stash
				meta.Put(stash.ID[:], serial)
				// Take care of diffing the stash
				stashDiffUpdate(oldSerial, stash, items[i], &stats, tx)
			}

			// Then update the metadata
			meta.Put(stash.ID[:], serial)

		}
		return nil
	})

}
