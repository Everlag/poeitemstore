package db

//go:generate msgp

import (
	"time"

	blake2b "github.com/minio/blake2b-simd"
)

// StringHeapID maps to a stored string identifier.
//
// This creates a layer of indirection when rebuilding items but
// saves on space for ids
type StringHeapID uint32

// StringHeapIDFromBytes generats the corresponding heap id
// from the provided bytes
func StringHeapIDFromBytes(bytes []byte) StringHeapID {
	return StringHeapID(btoi32(bytes))
}

// ToBytes returns the byte-wise represenation of a StringHeapID
func (id StringHeapID) ToBytes() []byte {
	return i32tob(uint32(id))
}

// LeagueHeapID maps to a stored string identifier specific to league
//
// This is basically StringHeapID but specialised for leagues
type LeagueHeapID uint16

// LeagueHeapIDFromSequence transforms a 64 bit bucket sequence number
// into a LeagueHeapID
func LeagueHeapIDFromSequence(seq uint64) LeagueHeapID {
	return LeagueHeapID(int16(seq))
}

// LeagueHeapIDFromBytes generats the corresponding heap id
// from the provided bytes
func LeagueHeapIDFromBytes(bytes []byte) LeagueHeapID {
	return LeagueHeapID(btoi16(bytes))
}

// ToBytes returns the byte-wise represenation of a LeagueHeapID
func (id LeagueHeapID) ToBytes() []byte {
	return i16tob(uint16(id))
}

// TimestampSize is the number of bytes used by Timestamp
//
// This is sized to minimize waste while
const TimestampSize = 4

// Timestamp is a compact represenation of a unix timestamp
type Timestamp [TimestampSize]byte

// NewTimestamp returns a Timestamp at the current time
func NewTimestamp() Timestamp {
	now := time.Now().Unix()

	nowBytes := i64tob(uint64(now))
	nowTrunc := nowBytes[TimestampSize:]

	var ts Timestamp
	copy(ts[:], nowTrunc)
	return ts
}

// ToTime converts a compact Timestamp to a time.Time
func (ts Timestamp) ToTime() time.Time {
	// Size the initial array with preceding zeroes
	fatBytes := make([]byte, 8-TimestampSize)
	// Jam the compact portion on
	fatBytes = append(fatBytes, ts[:]...)
	fatUint := btoi64(fatBytes)
	return time.Unix(int64(fatUint), 0)
}

// GGGIDSize is the size in bytes a derived ID can be
const GGGIDSize = 10

// GGGID is an Identifier derived from per-item/stash tab UID
//
// A PID is 80 bits = 10 bytes,
// this allows 2^40 items to be represented taking into birthdays
// and represents significant savings relative to the GGG api provided id
type GGGID [GGGIDSize]byte

// GGGIDFromUID generates an ID for internal use from a UID string
func GGGIDFromUID(uid string) GGGID {

	var id [GGGIDSize]byte

	hash := blake2b.Sum512([]byte(uid))

	copy(id[:], hash[:])

	return id
}

// IDSize is the size in bytes an internal identifier can be
const IDSize = 8

// ID is an Identifier calculated internally
//
// This is effectively just a 64 bit uint
type ID [IDSize]byte

// IDFromSequence converts a sequence number into an identifier
func IDFromSequence(seq uint64) ID {
	var id [IDSize]byte
	bin := i64tob(seq)
	copy(id[:], bin)

	return id
}

// ItemMod represents a compact explicit or implicit modifier on an item
//msgp:tuple ItemMod
type ItemMod struct {
	Mod    StringHeapID
	Values []uint16
}

// Item represents a compact record of an item.
//msgp:tuple Item
type Item struct {
	ID             ID
	GGGID          GGGID        // Allows mapping from simple ID to UUID
	Stash          GGGID        // Allows access to stash and corresponding metadata
	Name           StringHeapID // On StringHeap
	TypeLine       StringHeapID // On StringHeap
	Note           StringHeapID // On StringHeap
	RootType       StringHeapID // On StringHeap
	RootFlavor     StringHeapID // On StringHeap
	League         LeagueHeapID // On LeagueHeap
	Corrupted      bool
	Identified     bool
	Mods           []ItemMod
	When           Timestamp // When this stash update was processed
	UpdateSequence uint16    // The sequence number associated with this item
}

// Stash represents a compact record of a stash.
//msgp:tuple Stash
type Stash struct {
	ID          GGGID        // Reference value for this Stash
	AccountName string       // Account-wide name, we need nothing else to PM
	Items       []GGGID      // GGGIDs for all items stored in that Stash
	League      LeagueHeapID // LeagueHeapID as stashes are single-league
}

// Diff takes an older version of a Stash and determines which items,
// in terms of GGGID, need to be added and which need to be removed.
func (s Stash) Diff(old Stash) (add, remove []GGGID) {
	// Keep track of which items existed previously
	existing := make(map[GGGID]struct{})
	for _, id := range old.Items {
		existing[id] = struct{}{}
	}

	// Intersect the new Stash from the old
	add = make([]GGGID, 0)
	for _, id := range s.Items {
		// Check if item already exists, if it doesn't, we need to add it
		if _, ok := existing[id]; ok {
			// Remove it from the existing if it exist
			//
			// This will allow us to take the remaining items in existing
			// as those that are not found or shared in the new update and
			// then remove them.
			delete(existing, id)
		} else {
			// We need to add any items not found
			add = append(add, id)
		}
	}

	// Pull out all the remaining keys in existing to find the items that
	// are no longer present in the new update and need to be removed.
	remove = make([]GGGID, len(existing))[:0]
	for id := range existing {
		remove = append(remove, id)
	}

	return
}
