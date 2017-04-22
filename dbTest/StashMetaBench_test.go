package dbTest

import (
	"testing"

	"github.com/Everlag/poeitemstore/db"
)

// BenchmarkCompactAddStashesFast determines how fast adding a single
// response to the database is with an already initialized StringHeap
func BenchmarkCompactAddStashesFast(b *testing.B) {

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

	// Define when so we can change over time
	when := TimeOfStart

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Create a database and populate StringHeap, we want to exclude
		// the time this takes from the benchmark
		b.StopTimer()
		bdb := NewTempDatabase(b)
		_, _, err := db.StashStashToCompact(resp.Stashes, when, bdb)
		if err != nil {
			b.Fatalf("failed to convert fat stashes to compact, err=%s\n", err)
		}
		b.StartTimer()

		cStashes, cItems, err := db.StashStashToCompact(resp.Stashes, when, bdb)
		if err != nil {
			b.Fatalf("failed to convert fat stashes to compact, err=%s\n", err)
		}

		_, err = db.AddStashes(cStashes, cItems, bdb)
		if err != nil {
			b.Fatalf("failed to AddStashes, err=%s", err)
		}

		when = when.Add(TestTimeDeltas)
	}

}

// BenchmarkCompactFast determines how fast compacting a response
// using an initialized StringHeap is.
//
// This should let us determine where improvements need to be made
// in order to help BenchmarkCompactAddStashesFast
func BenchmarkCompactFast(b *testing.B) {

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

	// Define when so we can change over time
	when := TimeOfStart

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Create a database and populate StringHeap, we want to exclude
		// the time this takes from the benchmark
		b.StopTimer()
		bdb := NewTempDatabase(b)
		_, _, err := db.StashStashToCompact(resp.Stashes, when, bdb)
		if err != nil {
			b.Fatalf("failed to convert fat stashes to compact, err=%s\n", err)
		}
		b.StartTimer()

		_, _, err = db.StashStashToCompact(resp.Stashes, when, bdb)
		if err != nil {
			b.Fatalf("failed to convert fat stashes to compact, err=%s\n", err)
		}

		when = when.Add(TestTimeDeltas)
	}

}

// BenchmarkCompactFast determines how fast adding a response to the database
// is with an initialized StringHeap as well as already compacted items.
//
//
// This should let us determine where improvements need to be made
// in order to help BenchmarkCompactAddStashesFast
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

	// Define when so we can change over time
	when := TimeOfStart

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Create a database and populate StringHeap, we want to exclude
		// the time this takes from the benchmark
		b.StopTimer()
		bdb := NewTempDatabase(b)
		cStashes, cItems, err := db.StashStashToCompact(resp.Stashes, when, bdb)
		if err != nil {
			b.Fatalf("failed to convert fat stashes to compact, err=%s\n", err)
		}
		b.StartTimer()

		_, err = db.AddStashes(cStashes, cItems, bdb)
		if err != nil {
			b.Fatalf("failed to AddStashes, err=%s", err)
		}

		when = when.Add(TestTimeDeltas)
	}

}
