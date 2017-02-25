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
	return StringHeapID(btoi64(bytes))
}

func StringHeapIDToBytes(id StringHeapID) []byte {
	return i64tob(uint64(id))
}

// LeagueHeapID maps to a stored string identifier specific to league
//
// This is basically StringHeapID but specialised for leagues
type LeagueHeapID uint16

// LeagueHeapIDFromBytes generats the corresponding heap id
// from the provided bytes
func LeagueHeapIDFromBytes(bytes []byte) LeagueHeapID {
	return LeagueHeapID(btoi16(bytes))
}

func LeagueHeapIDToBytes(id LeagueHeapID) []byte {
	return i64tob(uint64(id))
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

// ItemMod represents a compact explicit or implicit modifier on an item
//msgp:tuple ItemMod
type ItemMod struct {
	Mod    StringHeapID
	Values []int
}

// Item represents a compact record of an item.
//msgp:tuple Item
type Item struct {
	ID         ID
	Stash      ID           // Allows access to stash and corresponding metadata
	Name       StringHeapID // On StringHeap
	TypeLine   StringHeapID // On StringHeap
	Note       StringHeapID // On StringHeap
	RootType   StringHeapID // On StringHeap
	RootFlavor StringHeapID // On StringHeap
	League     LeagueHeapID // On LeagueHeap
	Corrupted  bool
	Identified bool
	Mods       []ItemMod
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
	mods := make([][]string, 0)
	for i, item := range items {
		names[i] = item.Name
		typeLines[i] = item.TypeLine
		notes[i] = item.Name
		leagues[i] = item.League
		rootTypes[i] = item.RootType
		rootFlavors[i] = item.RootFlavor

		itemMods := make([]string, 0)
		concatMods := item.GetMods()
		for _, mod := range concatMods {
			itemMods = append(itemMods, string(mod.Template))
		}
		mods = append(mods, itemMods)
	}

	// Prepare to enter everything on the StringHeap
	// This method results in a single disk write for everything
	// rather than if we were to use separate SetStrings calls
	var nameIds, typeLineIds, noteIds,
		rootTypeIds, rootFlavorIds []StringHeapID
	modIds := make([][]StringHeapID, len(mods))

	gen := func(index int) ([]string, []StringHeapID) {

		// Nasty, hardcoded switch... a better solution sometime :|
		switch index {
		case 0:
			nameIds = make([]StringHeapID, len(names))
			return names, nameIds
		case 1:
			typeLineIds = make([]StringHeapID, len(typeLines))
			return typeLines, typeLineIds
		case 2:
			noteIds = make([]StringHeapID, len(notes))
			return notes, noteIds
		case 3:
			rootFlavorIds = make([]StringHeapID, len(rootFlavors))
			return rootFlavors, rootFlavorIds
		case 4:
			rootTypeIds = make([]StringHeapID, len(rootTypes))
			return rootTypes, rootTypeIds
		}

		if index > 4 && (index-4) < len(mods) {
			m := mods[index-4]
			modIds[index-4] = make([]StringHeapID, len(m))
			return m, modIds[index-4]
		}

		return nil, nil
	}

	err := SetStringsCB(gen, db)
	if err != nil {
		return nil, fmt.Errorf("failed to set strings in StringHeap, err=%s", err)
	}

	leagueIds, err := SetLeagues(leagues, db)
	if err != nil {
		return nil,
			fmt.Errorf("failed to add leagues to LeagueHeap, err=%s", err)
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

		// And now the worst part, the item mods :|
		//
		// This is very sensitive and makes for nasty indexing logic...
		itemMods := make([]ItemMod, len(modIds[i]))
		concatMods := item.GetMods()
		for k, mod := range modIds[i] {
			itemMods[k] = ItemMod{
				Mod:    mod,
				Values: concatMods[k].Values,
			}
		}
		compact[i].Mods = itemMods

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
		return nil, nil, err
	}

	return compact, compactItems, nil

}
