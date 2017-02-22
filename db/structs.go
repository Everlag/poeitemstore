package db

//go:generate msgp

import (
	"crypto/sha512"

	"fmt"

	"github.com/Everlag/gothing/stash"
	"github.com/boltdb/bolt"
)

// StringHeapID maps to a stored string identifier.
//
// This creates a layer of indirection when rebuilding items but
// saves on space for ids
type StringHeapID uint64

// StringHeapIDFromBytes generats the corresponding heap id
// from the provided bytes
func StringHeapIDFromBytes(bytes []byte) StringHeapID {
	return StringHeapID(btoi(bytes))
}

func StringHeapIDToBytes(id StringHeapID) []byte {
	return itob(uint64(id))
}

// IDSize is the size in bytes a derived ID can be
const IDSize = 12

// ID is an Identifier derived from per-item/stash tab UID
//
// A PID is 96 bits = 12 bytes,
// this allows 2^48 items to be represented taking into birthdays
// and represents significant savings relative to the GGG api provided id
type ID [IDSize]byte

// IDFromUID generates an ID for internal use from a UID string
func IDFromUID(uid string) ID {

	var id [IDSize]byte

	hash := sha512.Sum512([]byte(uid))

	copy(id[:], hash[:])

	return id
}

// Item represents a compact record of an item.
type Item struct {
	ID         ID
	Stash      ID           // Allows access to stash and corresponding metadata
	Name       StringHeapID // On StringHeap
	TypeLine   StringHeapID // On StringHeap
	Note       StringHeapID // On StringHeap
	League     StringHeapID // On StringHeap
	RootType   StringHeapID // On StringHeap
	RootFlavor StringHeapID // On StringHeap
	Corrupted  bool
	Identified bool
}

// StashItemsToCompact converts fat Item records to their compact form
//
// This also ensures all strings present on that item will be available
// on the StringHeap
func StashItemsToCompact(items []stash.Item, db *bolt.DB) ([]Item, error) {

	// Extract everything we need to put onto the string heap
	names := make([]string, len(items))
	typeLines := make([]string, len(items))
	notes := make([]string, len(items))
	leagues := make([]string, len(items))
	rootTypes := make([]string, len(items))
	rootFlavors := make([]string, len(items))
	for i, item := range items {
		names[i] = item.Name
		typeLines[i] = item.TypeLine
		notes[i] = item.Name
		leagues[i] = item.League
		rootTypes[i] = item.RootType
		rootFlavors[i] = item.RootFlavor
	}

	// Insert onto StringHeap while fetching their identifiers
	nameIds, err := SetStrings(names, db)
	if err != nil {
		return nil,
			fmt.Errorf("failed to add names to StringHeap, err=%s", err)
	}
	typeLineIds, err := SetStrings(typeLines, db)
	if err != nil {
		return nil,
			fmt.Errorf("failed to add typeLines to StringHeap, err=%s", err)
	}
	noteIds, err := SetStrings(notes, db)
	if err != nil {
		return nil,
			fmt.Errorf("failed to add notes to StringHeap, err=%s", err)
	}
	leagueIds, err := SetStrings(leagues, db)
	if err != nil {
		return nil,
			fmt.Errorf("failed to add leagues to StringHeap, err=%s", err)
	}
	rootTypeIds, err := SetStrings(rootTypes, db)
	if err != nil {
		return nil,
			fmt.Errorf("failed to add rootTypes to StringHeap, err=%s", err)
	}
	rootFlavorIds, err := SetStrings(rootFlavors, db)
	if err != nil {
		return nil,
			fmt.Errorf("failed to add rootFlavor to StringHeap, err=%s", err)
	}

	// Build compact items from the ids
	compact := make([]Item, len(items))

	for i, item := range items {
		compact[i] = Item{
			ID:         IDFromUID(item.ID),
			Stash:      IDFromUID(item.StashID),
			Name:       nameIds[i],
			TypeLine:   typeLineIds[i],
			Note:       noteIds[i],
			League:     leagueIds[i],
			RootType:   rootTypeIds[i],
			RootFlavor: rootFlavorIds[i],
			Identified: item.Identified,
			Corrupted:  item.Corrupted,
		}
	}

	return compact, nil

}

// Stash represents a compact record of a stash.
type Stash struct {
	ID          ID     // Reference value for this Stash
	AccountName string // Account-wide name, we need nothing else to PM
}

// StashStashToCompact converts fat Item records to their compact form
// while also stripping items out in their compact form.
func StashStashToCompact(stashes []stash.Stash,
	db *bolt.DB) ([]Stash, []Item, error) {

	// Compact stashes and flatten items
	compact := make([]Stash, len(stashes))
	flatItems := make([]stash.Item, 0)
	for i, stash := range stashes {
		compact[i] = Stash{
			AccountName: stash.AccountName,
			ID:          IDFromUID(stash.ID),
		}

		flatItems = append(flatItems, stash.Items...)
	}

	compactItems, err := StashItemsToCompact(flatItems, db)
	if err != nil {
		err = fmt.Errorf("failed to compact items, err=%s", err)
	}

	return compact, compactItems, nil

}