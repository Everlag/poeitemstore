package dbTest

import (
	"testing"

	"time"

	"github.com/Everlag/poeitemstore/cmd"
	"github.com/Everlag/poeitemstore/db"
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

// Test as searching across multiple stash updates
func TestIndexQuery48UpdatesMovespeedFireResist(t *testing.T) {

	t.Parallel()

	// bdb := NewTempDatabase(t)

	// // Define our search up here, it will be constant for all of
	// // our sub-tests
	// search := cmd.MultiModSearch{
	// 	MaxDesired: 4,
	// 	RootType:   "Armour",
	// 	RootFlavor: "Boots",
	// 	League:     "Legacy",
	// 	Mods: []string{
	// 		"#% increased Movement Speed",
	// 		"+#% to Fire Resistance",
	// 	},
	// 	MinValues: []uint16{
	// 		24,
	// 		27,
	// 	},
	// }

	// Fetch the changes we need
	set := GetChangeSet("data/testSet - 48 updates.msgp", t)
	if len(set.Changes) != 48 {
		t.Fatalf("wrong number of changes, expected 48 got %d",
			len(set.Changes))
	}

	inverter := GetChangeSetInverter(set)

	for i, comp := range set.Changes {
		id := inverter[i]
		start := time.Now()
		resp, err := comp.Decompress()
		if err != nil {
			t.Fatalf("failed to decompress stash.Compressed, changeID=%s err=%s",
				id, err)
		}
		end := time.Now()
		t.Logf("%d stashes, took %s", len(resp.Stashes), end.Sub(start))
	}

}
