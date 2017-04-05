package dbTest

import (
	"io/ioutil"
	"os"
	"testing"

	"fmt"

	"sync"

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
func NewTempDatabase(t *testing.T) *bolt.DB {

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

// GetTestStashUpdate returns the []db.Stash kept in the dbTest directory
// and accessible using Asset() from go-bindata.
//
// This requires a boltdb.DB as it performs compaction rather than
// returning []stash.Stash, this means the return values can be
// directly used in db.AddStashes.
func GetTestStashUpdate(path string, bdb *bolt.DB,
	t *testing.T) ([]db.Stash, [][]db.Item) {
	raw, err := Asset(path)
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

// CompareStats tests the provided stats and fails if they are mismatched
func CompareStats(expected, got *db.StashUpdateStats, t *testing.T) {
	if err := expected.Compare(got); err != nil {
		t.Fatalf("%s\n%s", err, got)
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

	t.Run("", func(t *testing.T) {})

}
