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
	bdb *bolt.DB, t testing.TB) (db.ItemStoreQuery, db.LeagueHeapID) {

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

// compareQueryResultsToExpected does exactly what it says.
//
// It performs translation of provided ids to GGGIDs so desired query results
// can be determined before hand and the GGGID can be
// used to specify specific items
func compareQueryResultsToExpected(ids []db.ID, league db.LeagueHeapID,
	expected []db.GGGID,
	bdb *bolt.DB, t *testing.T) (missing int, found int, err error) {

	// Fetch the items so we can grab their GGGIDs
	foundItems := make([]db.Item, 0)
	// Also fetch expected items so we can show details if not found
	expectedItems := make(map[db.GGGID]db.Item, 0)
	err = bdb.View(func(tx *bolt.Tx) error {
		for _, id := range ids {
			item, err := db.GetItemByID(id, league, tx)
			if err != nil {
				return fmt.Errorf("failed to find item, err=%s", err)
			}
			foundItems = append(foundItems, item)
		}

		expectedIDs, err := db.GetGGGIDTranslations(expected, league, tx)
		if err != nil {
			return fmt.Errorf("failed to translate expected ID, err=%s", err)
		}
		for _, id := range expectedIDs {
			item, err := db.GetItemByID(id, league, tx)
			if err != nil {
				return fmt.Errorf("failed to find item, err=%s", err)
			}
			expectedItems[item.GGGID] = item
		}

		return nil
	})
	if err != nil {
		err = fmt.Errorf("failed to find queried item in database, err=%s",
			err)
		return
	}

	lookup := make(map[db.GGGID]struct{})
	for _, id := range expected {
		lookup[id] = struct{}{}
	}

	// Check to ensure the item we want is found
	for _, item := range foundItems {
		if _, ok := lookup[item.GGGID]; ok {
			delete(lookup, item.GGGID)
		}
	}

	if len(lookup) != 0 {
		t.Logf("found items, GGGIDs are")
		for _, item := range foundItems {
			t.Logf("	%#x", item.GGGID)
			t.Logf(item.Inflate(bdb).Name)
		}
		t.Logf("--missing items, GGGIDs are--")
		for id := range lookup {
			t.Logf("	%#x", id)
			t.Logf(expectedItems[id].Inflate(bdb).Name)
		}
		t.Logf("failed to find expected items")
		missing = len(lookup)
		found = len(foundItems)
		err = fmt.Errorf("%d items missing despite %d found",
			missing, found)
		return
	}

	found = len(expected)
	return
}

// CompareQueryResultsToExpected does exactly what it says.
//
// This is just a simple wrapper for compareQueryResultsToExpected
// which fails if we encounter a failure
func CompareQueryResultsToExpected(ids []db.ID, league db.LeagueHeapID,
	expected []db.GGGID,
	bdb *bolt.DB, t *testing.T) {

	_, _, err := compareQueryResultsToExpected(ids, league, expected,
		bdb, t)
	if err != nil {
		t.Fatal(err)
	}
}

// IDsToGGGID translates provided IDs to their GGGID form
func IDsToGGGID(ids []db.ID, league db.LeagueHeapID,
	bdb *bolt.DB, t testing.TB) []db.GGGID {

	gggIDs := make([]db.GGGID, 0)
	err := bdb.View(func(tx *bolt.Tx) error {
		for _, id := range ids {
			item, err := db.GetItemByID(id, league, tx)
			if err != nil {
				return fmt.Errorf("failed to find item, err=%s", err)
			}
			gggIDs = append(gggIDs, item.GGGID)
		}

		return nil
	})
	if err != nil {
		t.Fatalf("failed to find queried item in database")
	}

	return gggIDs
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
		stashes, items := GetTestStashUpdate("singleStash.json",
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
		stashes, items := GetTestStashUpdate("singleStash - 3ItemsAdded.json",
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

// Test as searching within a single stash and the case of
// no matches. This is a degenerate case.
func TestItemStoreQuerySingleStashFindNone(t *testing.T) {

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
			"+#% to Cold Resistance",
			"#% increased Rarity of Items found",
		},
		MinValues: []uint16{
			20,
			20,
			10,
			100,
			100,
		},
	}

	// Test to ensure a good baseline
	t.Run("Baseline", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("singleStash.json",
			bdb, t)

		// This needs to be done AFTER the database has been populated
		query, _ := MultiModSearchToItemStoreQuery(search, bdb, t)

		_, err := db.AddStashes(stashes, items, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}

		// Run the search and translate into items
		ids, err := query.Run(bdb)
		if err != nil {
			t.Fatalf("failed to run query, err=%s", err)
		}
		if len(ids) > 0 {
			t.Fatalf("found items when impossible")
		}
	})

}

// Test as searching within multiple stashes
func TestItemStoreQuery11StashesSingleMod(t *testing.T) {

	t.Parallel()

	bdb := NewTempDatabase(t)

	// Define our search up here, it will be constant for all of
	// our sub-tests
	search := cmd.MultiModSearch{
		MaxDesired: 4,
		RootType:   "Armour",
		RootFlavor: "Quiver",
		League:     "Legacy",
		Mods: []string{
			"#% increased Global Critical Strike Chance",
		},
		MinValues: []uint16{
			11,
		},
	}

	// Keep the items we expect here.
	//
	// This will have items added between sub-tests when the database
	// is being manipulated.
	expected := []db.GGGID{
		db.GGGIDFromUID("0eebcc5526479e9193dbbb67f576c51e4926b810feb89464f88404118bede4e8"),
		db.GGGIDFromUID("58a00354024211d4e9c0d29a6f66bd5b36369ebfbba6d438883f815456a85029"),
	}

	// Test to ensure a good baseline
	t.Run("Baseline", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("11Stashes.json",
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
		db.GGGIDFromUID("4462e11ca8fbeeab11d83c156d10d00df9c69d7a05aba42f33d8110f84160b67"),
		db.GGGIDFromUID("eb77c02e62c18a9ebaf3e48c19c3df6d931eb7afd2758ab1c626edd9b54f6450"))

	// Test to ensure the added item can be found
	t.Run("3StashesAdded", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("11Stashes - 3StashesAddedWith92Items.json",
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

// Test as searching within multiple stashes
func TestItemStoreQuery11StashesMultiModA(t *testing.T) {

	t.Parallel()

	bdb := NewTempDatabase(t)

	// Define our search up here, it will be constant for all of
	// our sub-tests
	search := cmd.MultiModSearch{
		MaxDesired: 4,
		RootType:   "Armour",
		RootFlavor: "Boots",
		League:     "Legacy",
		Mods: []string{
			"#% increased Movement Speed",
			"+#% to Fire Resistance",
		},
		MinValues: []uint16{
			24,
			27,
		},
	}

	// Keep the items we expect here.
	//
	// This will have items added between sub-tests when the database
	// is being manipulated.
	expected := []db.GGGID{
		db.GGGIDFromUID("4c7382ccaaa9afe4c49498a21d7a9882a8db57528a732a4de2a886b5003263fe"),
		db.GGGIDFromUID("6d428356ff76f580dd461e072a179604946eccab73d65a197ef87c3cfa3754bf"),
		db.GGGIDFromUID("baa2c993ed0f4d9de5f1bf1a46d6cb43f64fd4de79dcdb10a50f8a43cc104898"),
	}

	// Test to ensure a good baseline
	t.Run("Baseline", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("11Stashes.json",
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
		db.GGGIDFromUID("12f142a8ffdaaf754510f66a75b4e27d6c3a954e5af6b36eb9844b8ca8b2fca3"))

	// Test to ensure the added item can be found
	t.Run("3StashesAdded", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("11Stashes - 3StashesAddedWith92Items.json",
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

// Test as searching within multiple stashes
func TestItemStoreQuery11StashesMultiModB(t *testing.T) {

	t.Parallel()

	bdb := NewTempDatabase(t)

	// Define our search up here, it will be constant for all of
	// our sub-tests
	search := cmd.MultiModSearch{
		MaxDesired: 4,
		RootType:   "Armour",
		RootFlavor: "Helmet",
		League:     "Legacy",
		Mods: []string{
			"#% increased Stun and Block Recovery",
			"#% increased Energy Shield",
		},
		MinValues: []uint16{
			10,
			100,
		},
	}

	// Keep the items we expect here.
	//
	// This will have items added between sub-tests when the database
	// is being manipulated.
	expected := []db.GGGID{
		db.GGGIDFromUID("13aaba8c5df27e0b9864724b498e10d5a1971a9bd54792736710829071e42d94"),
	}

	// Test to ensure a good baseline
	t.Run("Baseline", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("11Stashes.json",
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
		db.GGGIDFromUID("711ee6ab3ea7ddea76f3382f5cb8b87b0a9333b75c83bc388df9e08cb5a3aa13"))

	// Test to ensure the added item can be found
	t.Run("3StashesAdded", func(t *testing.T) {
		stashes, items := GetTestStashUpdate("11Stashes - 3StashesAddedWith92Items.json",
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
