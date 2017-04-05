package dbTest

import (
	"testing"

	"fmt"

	"github.com/Everlag/poeitemstore/cmd"
	"github.com/Everlag/poeitemstore/db"
	"github.com/boltdb/bolt"
)

// MultiModSearchToItemStoreQuery converts a MultiModSearch
// into a ItemStoreQuery. It also returns the league because
// you usually need that...
//
// Yes, I did steal this right out of commands...
func MultiModSearchToItemStoreQuery(search cmd.MultiModSearch,
	bdb *bolt.DB, t *testing.T) (db.ItemStoreQuery, db.LeagueHeapID) {

	if len(search.MinValues) != len(search.Mods) {
		t.Fatalf("each mod must have a minvalue")
	}

	// Lookup the root, flavor, and mod
	strings := []string{search.RootType, search.RootFlavor}
	ids, err := db.GetStrings(strings, bdb)
	if err != nil {
		t.Fatalf("failed to fetch rootType or RootFlavor id, err=%s\n", err)
	}
	modIds, err := db.GetStrings(search.Mods, bdb)
	if err != nil {
		t.Fatalf("failed to fetch mod id, err=%s\n", err)
	}

	// And we we need to fetch the league
	leagueIDs, err := db.GetLeagues([]string{search.League}, bdb)
	if err != nil {
		t.Fatalf("failed to fetch league, err=%s\n", err)
	}

	return db.NewItemStoreQuery(ids[0], ids[1],
		modIds, search.MinValues, leagueIDs[0], search.MaxDesired), leagueIDs[0]

}

// CompareQueryResultsToExpected does exactly what it says.
//
// It performs translation of provided ids to GGGIDs so desired query results
// can be determined before hand and the GGGID can be
// used to specify specific items
func CompareQueryResultsToExpected(ids []db.ID, league db.LeagueHeapID,
	expected []db.GGGID,
	bdb *bolt.DB, t *testing.T) {

	// Fetch the items so we can grab their GGGIDs
	foundItems := make([]db.Item, 0)
	err := bdb.View(func(tx *bolt.Tx) error {
		for _, id := range ids {
			item, err := db.GetItemByID(id, league, tx)
			if err != nil {
				return fmt.Errorf("failed to find item, err=%s", err)
			}
			foundItems = append(foundItems, item)
		}

		return nil
	})
	if err != nil {
		t.Fatalf("failed to find queried item in database")
	}

	lookup := make(map[db.GGGID]struct{})
	for _, id := range expected {
		lookup[id] = struct{}{}
	}

	// Check to ensure the item we want is found
	numberNeeded := len(expected)
	for _, item := range foundItems {
		if _, ok := lookup[item.GGGID]; ok {
			numberNeeded--
		}
	}

	if numberNeeded != 0 {
		t.Fatalf("failed to find expected items for query; missing %d items despite %d found",
			numberNeeded, len(foundItems))
	}
}

// Test as searching within a single stash
func TestItemStoreQuerySingleStash(t *testing.T) {

	t.Parallel()

	bdb := NewTempDatabase(t)

	// Define our search up here, it will be constant for all of
	// our sub-tests
	search := cmd.MultiModSearch{
		MaxDesired: 2,
		RootType:   "Jewelry",
		RootFlavor: "Ring",
		League:     "Legacy",
		Mods: []string{
			"+# to Strength",
			"+# to Intelligence",
			"+# to maximum Energy Shield",
		},
		MinValues: []uint16{
			20,
			20,
			10,
		},
	}

	// Keep the items we expect here.
	//
	// This will have items added between sub-tests when the database
	// is being manipulated.
	expected := []db.GGGID{
		db.GGGIDFromUID("3d474bb6f4d2b3bf86c0911aac89b5c50bef1d556240f745936df3b7d78a1db1"),
	}

	// Test to ensure a good baseline
	t.Run("Baseline", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("testData/singleStash.json",
			bdb, t)

		// This needs to be done AFTER the database has been populated
		query, league := MultiModSearchToItemStoreQuery(search, bdb, t)

		_, err := db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}

		// Run the search and translate into items
		ids, err := query.Run(bdb)
		if err != nil {
			t.Fatalf("failed to run query, err=%s", err)
		}
		CompareQueryResultsToExpected(ids, league, expected, bdb, t)
	})

	// Add in the next item we expect to find following the next mutation
	expected = append(expected,
		db.GGGIDFromUID("0125dab1d32f9e28d5531900d0d654774e7d8fc1e26bc717ada8e49231990f61"))

	// Test to ensure the added item can be found
	t.Run("3ItemsAdded", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("testData/singleStash - 3ItemsAdded.json",
			bdb, t)

		// This needs to be done AFTER the database has been populated
		query, league := MultiModSearchToItemStoreQuery(search, bdb, t)

		_, err := db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}

		// Run the search and translate into items
		ids, err := query.Run(bdb)
		if err != nil {
			t.Fatalf("failed to run query, err=%s", err)
		}
		CompareQueryResultsToExpected(ids, league, expected, bdb, t)
	})

}
