package dbTest

import (
	"testing"

	"github.com/Everlag/poeitemstore/db"
)

// Test adding whole stashes
func Test11StashesAdd(t *testing.T) {

	t.Parallel()

	bdb := NewTempDatabase(t)

	// Keep track of our base statistics, we'll need these
	var baseStats *db.StashUpdateStats

	// Test to ensure we can handle a single update
	t.Run("Baseline", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("testData/11Stashes.json",
			bdb, t)

		var err error
		baseStats, err = db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}
	})

	t.Run("3StashesAdded", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("testData/11Stashes - 3StashesAddedWith92Items.json",
			bdb, t)

		stats, err := db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}

		expected := &db.StashUpdateStats{
			Added:  3,
			Intact: 11,
			Items: db.ItemUpdateStats{
				Added: 92,
				Kept:  763,
			},
		}
		CompareStats(expected, stats, t)
	})

}

// Test adding some but not all of the previous stashes
func Test11StashesPartial(t *testing.T) {

	t.Parallel()

	bdb := NewTempDatabase(t)

	// Keep track of our base statistics, we'll need these
	var baseStats *db.StashUpdateStats

	// Test to ensure we can handle a single update
	t.Run("Baseline", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("testData/11Stashes.json",
			bdb, t)

		var err error
		baseStats, err = db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}
	})

	t.Run("2StashesNotIncluded", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("testData/11Stashes - 2StashesRemovedWith169Items.json",
			bdb, t)

		stats, err := db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}

		// Nothing should change, Intact is 9 as we only consider
		// stashes in the update rather than ALL the stashes.
		expected := &db.StashUpdateStats{
			Intact: 9,
			Items: db.ItemUpdateStats{
				Kept: 709,
			},
		}
		CompareStats(expected, stats, t)
	})

}
