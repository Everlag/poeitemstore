package dbTest

import (
	"testing"

	"github.com/Everlag/gothing/db"
)

// Test additions to a single stash on a per-item level
func TestSingleStashAdd(t *testing.T) {

	t.Parallel()

	bdb := NewTempDatabase(t)

	// Keep track of our base statistics, we'll need these
	var baseStats *db.StashUpdateStats

	// Test to ensure we can handle a single update
	t.Run("Baseline", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("testData/singleStash.json",
			bdb, t)

		var err error
		baseStats, err = db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}
	})

	t.Run("3ItemsAdded", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("testData/singleStash - 3ItemsAdded.json",
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
		stashes, items := GetTestStashUpdate("testData/singleStash.json",
			bdb, t)

		var err error
		baseStats, err = db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}
	})

	t.Run("3ItemsRemoved", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("testData/singleStash - 3ItemsRemoved.json",
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
		stashes, items := GetTestStashUpdate("testData/singleStash.json",
			bdb, t)

		var err error
		baseStats, err = db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}
	})

	t.Run("3ItemsAdded", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("testData/singleStash - 3ItemsAdded.json",
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
		stashes, items := GetTestStashUpdate("testData/singleStash - 3ItemsRemoved.json",
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
