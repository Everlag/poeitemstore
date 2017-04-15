package dbTest

import (
	"testing"

	"github.com/Everlag/poeitemstore/db"
)

var benchQueryResult []db.ID

// BenchmarkSingleIndexQuery runs a single query on the database.
//
// Consider this the absolute best case scenario as the cache will
// be as hot as it can possibly be.
func BenchmarkSingleIndexQuery(b *testing.B) {

	b.ReportAllocs()

	bdb := NewTempDatabase(b)

	// Define our search up here, it will be constant for all of
	// our sub-tests
	search := QueryBootsMovespeedFireResist.Clone()

	// Fetch the changes we need
	set := GetChangeSet("testSet - 11 updates.msgp", b)
	if len(set.Changes) != 11 {
		b.Fatalf("wrong number of changes, expected 11 got %d",
			len(set.Changes))
	}

	// Run the changes, nop for the callback as
	// we only care about the end result
	RunChangeSet(set, func(id string) error {
		return nil
	}, bdb, b)

	// Translate the query now, after we are more likely
	// to have the desired mods available on the StringHeap
	indexQuery, _ := MultiModSearchToIndexQuery(search, bdb, b)

	b.ResetTimer()

	var err error

	for i := 0; i < b.N; i++ {
		benchQueryResult, err = indexQuery.Run(bdb)
		if err != nil {
			b.Fatalf("failed IndexQuery.Run, err=%s", err)
		}

		if len(benchQueryResult) < search.MaxDesired {
			b.Fatalf("failed to find enough results in query")
		}
	}
}
