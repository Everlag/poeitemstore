package dbTest

import (
	"testing"

	"time"

	"github.com/Everlag/poeitemstore/cmd"
	"github.com/Everlag/poeitemstore/db"
	"github.com/boltdb/bolt"
)

var benchQueryResult []db.ID

// setupBenchDB prepares a database with a ChangeSet located at path
// and returns it.
//
// NOTE: we use a different time delta here to better represent
// real operations of the index.
func setupBenchDB(path string, b *testing.B) *bolt.DB {
	bdb := NewTempDatabase(b)

	// Fetch the changes we need
	set := GetChangeSet(path, b)

	// Run the changes, nop for the callback as
	// we only care about the end result
	RunChangeSet(set, func(id string) error {
		return nil
	}, TimeOfStart, time.Second*20, bdb, b)

	return bdb
}

// runBenchQuery runs a provided search in the context of a benchmark
//
// This will likely incur overhead but thats a static cost we eat.
func runBenchQuery(search cmd.MultiModSearch, bdb *bolt.DB, b *testing.B) {
	// Translate the query now, after we are more likely
	// to have the desired mods available on the StringHeap.
	//
	// This is done within the benchmarking time as it must be done
	// for any query
	indexQuery, _ := MultiModSearchToIndexQuery(search, bdb, b)

	benchQueryResult, err := indexQuery.Run(bdb)
	if err != nil {
		b.Fatalf("failed IndexQuery.Run, err=%s", err)
	}

	// Sanity check
	if len(benchQueryResult) < search.MaxDesired {
		b.Fatalf("failed to find enough results in query")
	}
}

// BenchmarkSingleIndexQuery runs a single query on the database.
//
// Consider this the absolute best case scenario as the cache will
// be as hot as it can possibly be.
func BenchmarkSingleIndexQueryFast(b *testing.B) {

	bdb := setupBenchDB("testSet - 11 updates.msgp", b)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runBenchQuery(QueryBootsMovespeedFireResist, bdb, b)
	}
}

// BenchmarkFiveIndexQuery runs five queries on the database.
//
// This should, hopefully, touch enough of the database to overflow the cache.
func BenchmarkFiveIndexQueryFast(b *testing.B) {

	bdb := setupBenchDB("testSet - 11 updates.msgp", b)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runBenchQuery(QueryBootsMovespeedFireResist.Clone(), bdb, b)
		runBenchQuery(QueryAmuletColdCritMulti.Clone(), bdb, b)
		runBenchQuery(QueryRingStrengthIntES.Clone(), bdb, b)
		runBenchQuery(QueryQuiverCritChance.Clone(), bdb, b)
		runBenchQuery(QueryHelmetRecoveryES.Clone(), bdb, b)
	}
}

// BenchmarkInterleavedLeagueIndexQuery runs ten queries on the database,
// each one alternating in league.
//
// This should, hopefully, touch enough of the database to overflow the cache.
func BenchmarkInterleavedLeagueIndexQueryFast(b *testing.B) {

	bdb := setupBenchDB("testSet - 11 updates.msgp", b)

	queries := []cmd.MultiModSearch{
		QueryBootsMovespeedFireResist,
		QueryAmuletColdCritMulti,
		QueryRingStrengthIntES,
		QueryQuiverCritChance,
		QueryHelmetRecoveryES,
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		for _, q := range queries {
			search := q.Clone()
			runBenchQuery(search, bdb, b)
			// Stash updates should have a mix of permanent and temp leagues
			search.League = "Standard"
			runBenchQuery(search, bdb, b)
		}
	}
}

// BenchmarkSingleIndexQuery runs a single query on the database.
//
// Consider this the absolute best case scenario as the cache will
// be as hot as it can possibly be.
func BenchmarkSingleIndexQuerySlow(b *testing.B) {

	bdb := setupBenchDB("testSet - 140 updates.msgp", b)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runBenchQuery(QueryBootsMovespeedFireResist, bdb, b)
	}
}

// BenchmarkFiveIndexQuery runs five queries on the database.
//
// This should, hopefully, touch enough of the database to overflow the cache.
func BenchmarkFiveIndexQuerySlow(b *testing.B) {

	bdb := setupBenchDB("testSet - 140 updates.msgp", b)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runBenchQuery(QueryBootsMovespeedFireResist.Clone(), bdb, b)
		runBenchQuery(QueryAmuletColdCritMulti.Clone(), bdb, b)
		runBenchQuery(QueryRingStrengthIntES.Clone(), bdb, b)
		runBenchQuery(QueryQuiverCritChance.Clone(), bdb, b)
		runBenchQuery(QueryHelmetRecoveryES.Clone(), bdb, b)
	}
}

// BenchmarkInterleavedLeagueIndexQuery runs ten queries on the database,
// each one alternating in league.
//
// This should, hopefully, touch enough of the database to overflow the cache.
func BenchmarkInterleavedLeagueIndexQuerySlow(b *testing.B) {

	bdb := setupBenchDB("testSet - 140 updates.msgp", b)

	queries := []cmd.MultiModSearch{
		QueryBootsMovespeedFireResist,
		QueryAmuletColdCritMulti,
		QueryRingStrengthIntES,
		QueryQuiverCritChance,
		QueryHelmetRecoveryES,
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		for _, q := range queries {
			search := q.Clone()
			runBenchQuery(search, bdb, b)
			// Stash updates should have a mix of permanent and temp leagues
			search.League = "Standard"
			runBenchQuery(search, bdb, b)
		}
	}
}
