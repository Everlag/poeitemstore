package dbTest

import (
	"testing"

	"github.com/Everlag/poeitemstore/db"
)

// BenchmarkAddStashesNoStringHeap determines how fast adding a single
// response to the database is with an uninitialized StringHeap.
func BenchmarkAddStashesNoStringHeapFast(b *testing.B) {

	// Fetch the changes we need
	set := GetChangeSet("testSet - 11 updates.msgp", b)
	change, err := set.GetFirstChange()
	if err != nil {
		b.Fatalf("failed to get first change")
	}

	resp, err := change.Decompress()
	if err != nil {
		b.Fatalf("failed to decompress response")
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Create a database, we want to exclude
		// the time this takes from the benchmark
		b.StopTimer()
		bdb := NewTempDatabase(b)
		b.StartTimer()

		cStashes, cItems, err := db.StashStashToCompact(resp.Stashes, bdb)
		if err != nil {
			b.Fatalf("failed to convert fat stashes to compact, err=%s\n", err)
		}

		_, err = db.AddStashes(cStashes, cItems, bdb)
		if err != nil {
			b.Fatalf("failed to AddStashes, err=%s", err)
		}
	}

}

// BenchmarkAddStashes determines how fast adding a single
// response to the database is with an already initialized StringHeap
func BenchmarkAddStashesFast(b *testing.B) {

	// Fetch the changes we need
	set := GetChangeSet("testSet - 11 updates.msgp", b)
	change, err := set.GetFirstChange()
	if err != nil {
		b.Fatalf("failed to get first change")
	}

	resp, err := change.Decompress()
	if err != nil {
		b.Fatalf("failed to decompress response")
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Create a database and populate StringHeap, we want to exclude
		// the time this takes from the benchmark
		b.StopTimer()
		bdb := NewTempDatabase(b)
		_, _, err := db.StashStashToCompact(resp.Stashes, bdb)
		if err != nil {
			b.Fatalf("failed to convert fat stashes to compact, err=%s\n", err)
		}
		b.StartTimer()

		cStashes, cItems, err := db.StashStashToCompact(resp.Stashes, bdb)
		if err != nil {
			b.Fatalf("failed to convert fat stashes to compact, err=%s\n", err)
		}

		_, err = db.AddStashes(cStashes, cItems, bdb)
		if err != nil {
			b.Fatalf("failed to AddStashes, err=%s", err)
		}
	}

}
