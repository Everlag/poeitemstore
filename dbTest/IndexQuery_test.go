package dbTest

import (
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
	bdb *bolt.DB, t testing.TB) (db.IndexQuery, db.LeagueHeapID) {

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

// IndexQueryWithResultsToItemStoreQuery converts a MultiModSearch
// into an ItemStoreQuery while attempting to preserve the semantics
// of an IndexQuery in the resulting ItemStoreQuery
//
// IndexQuery has results ordered by highest values
// while ItemStoreQuery has results ordered by latest additions
// with minimum values.
func IndexQueryWithResultsToItemStoreQuery(search cmd.MultiModSearch,
	prevResults []stash.Item,
	bdb *bolt.DB, t testing.TB) db.ItemStoreQuery {

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
	// Pre-populate with pre-existing values, found items will
	// always be equal to or higher than the pre-existing
	for i, mod := range search.Mods {
		minValueMap[mod] = search.MinValues[i]
	}
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

// ChangeSetUse is the callback given to RunChangeSet
// to make traversing a ChangeSet less awful.
//
// ChangeSetUse is expected to be an anonymous function
// accessing the database through its defining scope.
type ChangeSetUse func(id string) error

// RunChangeSet steps through a given ChangeSet, adding changes
// to the provided DB then calling cb to do some work
// on the database.
//
// cb will we called for each entry in the ChangeSet
func RunChangeSet(set stash.ChangeSet, cb ChangeSetUse,
	bdb *bolt.DB, t testing.TB) {

	// Generate a mapping of change to id we'll need
	inverter := GetChangeSetInverter(set)

	for i, comp := range set.Changes {
		// Decompress
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

		if err := cb(id); err != nil {
			t.Fatalf("failed to cb in RunChangeSet, err=%s", err)
		}
	}

}

var QueryBootsMovespeedFireResist = cmd.MultiModSearch{
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

var QueryAmuletColdCritMulti = cmd.MultiModSearch{
	MaxDesired: 4,
	RootType:   "Jewelry",
	RootFlavor: "Amulet",
	League:     "Legacy",
	Mods: []string{
		"#% increased Cold Damage",
		"+#% to Global Critical Strike Multiplier",
	},
	MinValues: []uint16{
		10,
		10,
	},
}

// testIndexQueryAgainstChangeSet ensures a given MultiModSearch
// is valid for every change in the ChangeSet located at path
func testIndexQueryAgainstChangeSet(search cmd.MultiModSearch, path string,
	t *testing.T) {

	t.Parallel()

	bdb := NewTempDatabase(t)

	// Fetch the changes we need
	set := GetChangeSet(path, t)
	if len(set.Changes) != 11 {
		t.Fatalf("wrong number of changes, expected 11 got %d",
			len(set.Changes))
	}

	// We have to find items that match at least once or else the test
	// is absolutely useless.
	foundOnce := false

	RunChangeSet(set, func(id string) error {
		// Translate the query now, after we are more likely
		// to have the desired mods available on the StringHeap
		indexQuery, league := MultiModSearchToIndexQuery(search, bdb, t)

		indexResult, err := indexQuery.Run(bdb)
		if err != nil {
			t.Fatalf("failed IndexQuery.Run, err=%s", err)
		}

		foundOnce = foundOnce || (len(indexResult) > 0)
		if len(indexResult) > 0 {
			t.Logf("found %d items", len(indexResult))
		}

		// Ensure correctness
		CompareIndexQueryResultsToItemStoreEquiv(search, indexResult, league,
			bdb, t)
		return nil
	}, bdb, t)

	if !foundOnce {
		t.Fatalf("failed to match any items across all queries")
	}
}

// Test as searching across multiple stash updates
func TestIndexQuery11UpdatesMovespeedFireResist(t *testing.T) {
	testIndexQueryAgainstChangeSet(QueryBootsMovespeedFireResist.Clone(),
		"testSet - 11 updates.msgp", t)
}

// Test as searching across multiple stash updates
func TestIndexQuery11UpdatesColdCritMulti(t *testing.T) {
	testIndexQueryAgainstChangeSet(QueryAmuletColdCritMulti.Clone(),
		"testSet - 11 updates.msgp", t)
}
