package dbTest

import (
	"testing"

	"github.com/Everlag/poeitemstore/db"
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

// Test additions to a single stash on a per-item level
func TestSingleStashAdd(t *testing.T) {

	t.Parallel()

	bdb := NewTempDatabase(t)

	// Keep track of our base statistics, we'll need these
	var baseStats *db.StashUpdateStats

	// Test to ensure we can handle a single update
	t.Run("Baseline", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("singleStash.json",
			bdb, t)

		var err error
		baseStats, err = db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}
	})

	t.Run("3ItemsAdded", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("singleStash - 3ItemsAdded.json",
			bdb, t)

		stats, err := db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}

		expected := &db.StashUpdateStats{
			Updated: 1,
			Items: db.ItemUpdateStats{
				Added: 3,
				Kept:  54,
			},
		}
		CompareStats(expected, stats, t)
	})

}

// Test additions to a single stash on a per-item level
func TestSingleStashRemove(t *testing.T) {

	t.Parallel()

	bdb := NewTempDatabase(t)

	// Keep track of our base statistics, we'll need these
	var baseStats *db.StashUpdateStats

	// Test to ensure we can handle a single update
	t.Run("Baseline", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("singleStash.json",
			bdb, t)

		var err error
		baseStats, err = db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}
	})

	t.Run("3ItemsRemoved", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("singleStash - 3ItemsRemoved.json",
			bdb, t)

		stats, err := db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}

		expected := &db.StashUpdateStats{
			Updated: 1,
			Items: db.ItemUpdateStats{
				Removed: 3,
				Kept:    48,
			},
		}
		CompareStats(expected, stats, t)
	})

}

// Test additions and removals for a single stash on a per-item level
func TestSingleStashAddAndRemove(t *testing.T) {

	t.Parallel()

	bdb := NewTempDatabase(t)

	// Keep track of our base statistics, we'll need these
	var baseStats *db.StashUpdateStats

	// Test to ensure we can handle a single update
	t.Run("Baseline", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("singleStash.json",
			bdb, t)

		var err error
		baseStats, err = db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}
	})

	t.Run("3ItemsAdded", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("singleStash - 3ItemsAdded.json",
			bdb, t)

		stats, err := db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}

		expected := &db.StashUpdateStats{
			Updated: 1,
			Items: db.ItemUpdateStats{
				Added: 3,
				Kept:  54,
			},
		}
		CompareStats(expected, stats, t)
	})

	t.Run("3ItemsRemoved", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("singleStash - 3ItemsRemoved.json",
			bdb, t)

		stats, err := db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}

		// NOTE: since this is run AFTER 3ItemsAdded, we instead have
		// 6 items removed and 45 kept
		expected := &db.StashUpdateStats{
			Updated: 1,
			Items: db.ItemUpdateStats{
				Removed: 6,
				Kept:    45,
			},
		}
		CompareStats(expected, stats, t)
	})
}

// Test fetching an item out of the singleStash and ensure it
// is exactly equivalent to what we expect.
func TestSingleStashUnmodified(t *testing.T) {

	t.Parallel()

	bdb := NewTempDatabase(t)

	// Test to ensure we can handle a single update
	stashes, items := GetTestStashUpdate("singleStash.json",
		bdb, t)

	// Add the items
	var err error
	_, err = db.AddStashes(stashes, items, bdb)
	if err != nil {
		t.Fatalf("failed to AddStashes, err=%s", err)
	}

	// Ensure each item we added is equivalent as we added it
	err = bdb.View(func(tx *bolt.Tx) error {
		for _, stashLevel := range items {
			for _, item := range stashLevel {
				id, err := db.GetGGGIDTranslations([]db.GGGID{item.GGGID},
					item.League, tx)
				if err != nil {
					return errors.Wrap(err, "failed to translate item")
				}

				found, err := db.GetItemByID(id[0], item.League, tx)
				if err != nil {
					return errors.Wrap(err, "failed to get item after translation")
				}

				// TODO: equality check of found vs item
				if !item.Equal(found) {
					t.Log("expected", item)
					t.Log("found", found)
					t.Fatalf("mismatched items when retrieving")
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("failed to validate items are same as entered, err=%s", err)
	}
}
