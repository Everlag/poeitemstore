package dbTest

import (
	"io/ioutil"
	"os"
	"testing"

	"fmt"

	"sync"

	"path/filepath"

	"github.com/Everlag/poeitemstore/cmd"
	"github.com/Everlag/poeitemstore/db"
	"github.com/Everlag/poeitemstore/stash"
	"github.com/boltdb/bolt"
)

// TempEnviron represents the complete environment for
// a test which must be removed during cleanup
type TempEnviron struct {
	DB   *bolt.DB
	File *os.File
}

// Close deletes all state associated with a TempEnviron
func (env *TempEnviron) Close() {
	if err := env.DB.Close(); err != nil {
		fmt.Printf("failed to close boltdb database, path=%s err=%s\n",
			env.File.Name(), err)
		return
	}

	if err := env.File.Close(); err != nil {
		fmt.Printf("failed to close testing file, path=%s err=%s\n",
			env.File.Name(), err)
		return
	}

	if err := os.Remove(env.File.Name()); err != nil {
		fmt.Printf("failed to remove testing file, path=%s err=%s\n",
			env.File.Name(), err)
		return
	}
}

// All of the testing environments we have accompanied
// by the lock used to prevent concurrent modification
var environments []TempEnviron
var envSync sync.Mutex

// NewTempDatabase prepares a fresh boltdb database for testing
// which will automatically be cleaned up after testing
func NewTempDatabase(t testing.TB) *bolt.DB {

	f, err := ioutil.TempFile("", "gothingTest")
	if err != nil {
		t.Fatalf("failed to open TempFile, err=%s", err)
	}

	db, err := db.Boot(f.Name())
	if err != nil {
		t.Fatalf("failed to open db, err=%s", err)
	}

	// Record our state for later removal in a safe way
	envSync.Lock()
	environments = append(environments, TempEnviron{db, f})
	defer envSync.Unlock()

	return db
}

// GetTestData returns the content of a file in testdata
//
// Provided path is package-relative,
// ie 'thing.json' fetches 'testdata/thing.json'
func GetTestData(path string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join("testdata", path))
}

// GetTestStashUpdate returns the []db.Stash kept in the dbTest directory
// and accessible using Asset() from go-bindata.
//
// This requires a boltdb.DB as it performs compaction rather than
// returning []stash.Stash, this means the return values can be
// directly used in db.AddStashes.
func GetTestStashUpdate(path string, bdb *bolt.DB,
	t *testing.T) ([]db.Stash, [][]db.Item) {
	raw, err := GetTestData(path)
	if err != nil {
		t.Fatalf("failed to fetch '%s', err=%s", path, err)
	}

	resp, err := stash.RespFromJSON(raw)
	if err != nil {
		t.Fatalf("failed to unmarshal testing json, path=%s, err=%s",
			path, err)
	}

	cStashes, cItems, err := db.StashStashToCompact(resp.Stashes, bdb)
	if err != nil {
		t.Fatalf("failed to convert fat stashes to compact, err=%s\n", err)
	}

	return cStashes, cItems
}

// GetChangeSet returns a stash.ChangeSet serialized at the provided path
// which is accessible using Assets() from go-bindata
func GetChangeSet(path string, t testing.TB) stash.ChangeSet {

	raw, err := GetTestData(path)
	if err != nil {
		t.Fatalf("failed to fetch '%s', err=%s", path, err)
	}

	var set stash.ChangeSet
	if _, err := set.UnmarshalMsg(raw); err != nil {
		t.Fatalf("failed to unmarshal '%s', err=%s",
			path, err)
	}

	return set
}

// GetChangeSetInverter returns a mapping from index in a ChangeSet
// to the ChangeID associated with that stash update
func GetChangeSetInverter(set stash.ChangeSet) map[int]string {
	inverter := make(map[int]string)
	for k, v := range set.ChangeIDToIndex {
		inverter[v] = k
	}
	return inverter
}

// QueryResultsToItems converts the provided IDs to their
// inflated stash forms
func QueryResultsToItems(ids []db.ID, league db.LeagueHeapID,
	bdb *bolt.DB, t testing.TB) []stash.Item {

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

// CompareStats tests the provided stats and fails if they are mismatched
func CompareStats(expected, got *db.StashUpdateStats, t testing.TB) {
	if err := expected.Compare(got); err != nil {
		t.Fatalf("%s\n%s", err, got)
	}
}

// CompareIndexQueryResultsToItemStoreEquiv ensures the correctness
// of results found with an IndexQuery by generating the equivalent
// ItemStoreQuery and testing against the results found there.
//
// NOTE: given the differing semantics between an IndexQuery
// and an ItemStoreQuery which cannot be completely mitigated by translation,
// this is a relaxed rather than absolute comparison.
func CompareIndexQueryResultsToItemStoreEquiv(search cmd.MultiModSearch,
	indexResult []db.ID, league db.LeagueHeapID,
	bdb *bolt.DB, t *testing.T) {
	// Inflate items for easier mod checking
	indexItems := QueryResultsToItems(indexResult, league, bdb, t)

	// Convert to equivalent ItemStoreQuery
	itemStoreQuery := IndexQueryWithResultsToItemStoreQuery(search,
		indexItems, bdb, t)

	// Perform the ItemStoreQuery to generate reference
	itemStoreResults, err := itemStoreQuery.Run(bdb)
	if err != nil {
		t.Fatalf("failed ItemStoreQuery.Run, err=%s", err)
	}

	// Translate :|
	itemStoreResultsGGG := IDsToGGGID(itemStoreResults, league, bdb, t)

	// Ensure they match
	missing, found, err := compareQueryResultsToExpected(indexResult, league,
		itemStoreResultsGGG, bdb, t)
	// Failure for exact match is handled leniently
	if err != nil {
		// We require at least once matching item between them
		if found < 1 {
			t.Fatal(err)
		} else {
			t.Logf("comparison allowing %d of %d items to be missing",
				missing, found)
		}

		// All items must at least satisfy the query to be comfortable
		if !search.Satisfies(indexItems) {
			t.Fatalf("indexResults do not satisfy MultiModSearch")
		}
	}
}

// TestMain prepares tests to be run
func TestMain(m *testing.M) {

	// Prep environment
	environments = make([]TempEnviron, 0)

	ret := m.Run()

	// Remove all environments
	for _, env := range environments {
		env.Close()
	}

	os.Exit(ret)
}

// Consider this a template of
func TestApples(t *testing.T) {

	t.Parallel()

	db := NewTempDatabase(t)
	fmt.Printf("I have a database, neat!, db=%s\n", db)

	t.Run("", func(t *testing.T) {})

}
