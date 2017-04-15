package dbTest

import (
	"fmt"
	"testing"

	"github.com/Everlag/poeitemstore/cmd"
	"github.com/Everlag/poeitemstore/db"
	"github.com/Everlag/poeitemstore/stash"
	"github.com/boltdb/bolt"
)

// MultiModSearchToIndexQuery converts a MultiModSearch
// into an IndexQuery. It also returns the league because
// you usually need that...
func MultiModSearchToIndexQuery(search cmd.MultiModSearch,
	bdb *bolt.DB, t *testing.T) (db.IndexQuery, db.LeagueHeapID) {

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

	return db.NewIndexQuery(ids[0], ids[1],
		modIds, search.MinValues, leagueIDs[0], search.MaxDesired), leagueIDs[0]

}

// QueryResultsToItems converts the provided IDs to their
// inflated stash forms
func QueryResultsToItems(ids []db.ID, league db.LeagueHeapID,
	bdb *bolt.DB, t *testing.T) []stash.Item {

	// Fetch the items so we can grab their GGGIDs
	compact := make([]db.Item, 0)
	err := bdb.View(func(tx *bolt.Tx) error {
		for _, id := range ids {
			item, err := db.GetItemByID(id, league, tx)
			if err != nil {
				return fmt.Errorf("failed to find item, err=%s", err)
			}
			compact = append(compact, item)
		}

		return nil
	})
	if err != nil {
		t.Fatalf("failed to find queried item in database")
	}
	foundItems := make([]stash.Item, 0)
	for _, tiny := range compact {
		foundItems = append(foundItems, tiny.Inflate(bdb))
	}

	return foundItems
}

// IndexQueryWithResultsToItemStoreQuery converts a MultiModSearch
// into an ItemStoreQuery while attempting to preserve the semantics
// of an IndexQuery in the resulting ItemStoreQuery
//
// IndexQuery has results ordered by highest values
// while ItemStoreQuery has results ordered by latest additions
// with minimum values.
func IndexQueryWithResultsToItemStoreQuery(search cmd.MultiModSearch,
	prevResults []stash.Item,
	bdb *bolt.DB, t *testing.T) db.ItemStoreQuery {

	if len(search.MinValues) != len(search.Mods) {
		t.Fatalf("each mod must have a minvalue")
	}

	// Setup a interestedMap so we can have constant time lookup
	// for which mods we are interesed in
	interestedMap := make(map[string]struct{})
	for _, mod := range search.Mods {
		interestedMap[mod] = struct{}{}
	}

	// Setup the minValue map, this will determine the real minimum
	// values which the ItemStoreQuery will need to find
	minValueMap := make(map[string]uint16)
	for _, item := range prevResults {
		for _, mod := range item.GetMods() {
			// Check if we are about this mod
			_, ok := interestedMap[string(mod.Template)]
			if !ok {
				continue
			}

			// Update the minValues as necessary
			prev, ok := minValueMap[string(mod.Template)]
			if !ok {
				prev = mod.Values[0]
			}
			if prev >= mod.Values[0] {
				minValueMap[string(mod.Template)] = mod.Values[0]
			}
		}
	}

	// Overwrite the search with the new minimum values
	prevLength := len(search.Mods) // Store old length for later
	search.Mods = make([]string, 0)
	search.MinValues = make([]uint16, 0)
	for mod, min := range minValueMap {
		search.Mods = append(search.Mods, mod)
		search.MinValues = append(search.MinValues, min)
	}
	if len(search.Mods) != prevLength {
		t.Fatalf("bad MultiModSearch translation: mismatched #mods")
	}

	t.Logf("Generated MultiModSearch:\n %s", search.String())

	itemStoreSearch, _ := MultiModSearchToItemStoreQuery(search, bdb, t)
	return itemStoreSearch

}

// Test as searching across multiple stash updates
func TestIndexQuery48UpdatesMovespeedFireResist(t *testing.T) {

	// t.Parallel()

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

	// Fetch the changes we need
	set := GetChangeSet("testSet - 11 updates.msgp", t)
	if len(set.Changes) != 11 {
		t.Fatalf("wrong number of changes, expected 11 got %d",
			len(set.Changes))
	}

	inverter := GetChangeSetInverter(set)

	for i, comp := range set.Changes {
		id := inverter[i]
		resp, err := comp.Decompress()
		if err != nil {
			t.Fatalf("failed to decompress stash.Compressed, changeID=%s err=%s",
				id, err)
		}

		t.Logf("processing changeID=%s", id)

		cStashes, cItems, err := db.StashStashToCompact(resp.Stashes, bdb)
		if err != nil {
			t.Fatalf("failed to convert fat stashes to compact, err=%s\n", err)
		}

		_, err = db.AddStashes(cStashes, cItems, bdb)
		if err != nil {
			t.Fatalf("failed to AddStashes, err=%s", err)
		}

		// Translate the query now, after we are more likely
		// to have the desired mods available on the StringHeap
		indexQuery, league := MultiModSearchToIndexQuery(search, bdb, t)

		indexResult, err := indexQuery.Run(bdb)
		if err != nil {
			t.Fatalf("failed IndexQuery.Run, err=%s", err)
		}

		// Ensure correctness
		CompareIndexQueryResultsToItemStoreEquiv(search, indexResult, league,
			bdb, t)

	}

}
