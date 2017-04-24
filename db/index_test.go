package db

import "testing"
import "bytes"

func TestIndexEntryAppendGet(t *testing.T) {
	ids := []ID{
		IDFromSequence(1),
		IDFromSequence(2),
		IDFromSequence(3),
		IDFromSequence(4),
		IDFromSequence(500),
		IDFromSequence(20),
		IDFromSequence(19),
		IDFromSequence(18),
	}

	entry := IndexEntry(nil)
	for _, id := range ids {
		entry = IndexEntryAppend(entry, id)
	}

	tinyIDs := entry.GetIDs(nil)

	if len(tinyIDs) != len(ids) {
		t.Fatalf("mismatched lengths, %d decompressed != %d original ids",
			len(tinyIDs), len(ids))
	}

	for i, id := range ids {
		if !bytes.Equal(id[:], tinyIDs[i][:]) {
			t.Fatal("mismatched compressed and decompressed results")
		}
	}
}
