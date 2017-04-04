package dbTest

import (
	"io/ioutil"
	"os"
	"testing"

	"fmt"

	"sync"

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

	db, err := bolt.Open(f.Name(), 0777, nil)
	if err != nil {
		t.Fatalf("failed to open db, err=%s", err)
	}

	// Record our state for later removal in a safe way
	envSync.Lock()
	environments = append(environments, TempEnviron{db, f})
	defer envSync.Unlock()

	return db
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

func TestApples(t *testing.T) {

	db := NewTempDatabase(t)
	fmt.Printf("I have a database, neat!, db=%s\n", db)

	t.Run("", func(t *testing.T) {})

	t.Run("", func(t *testing.T) {})

}
