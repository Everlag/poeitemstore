package db

import (
	"time"

	"github.com/Everlag/poeitemstore/stash"
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

// Inflate returns the string represented by the given StringHeapID
func (id StringHeapID) Inflate(db *bolt.DB) string {
	return InflateString(id, db)
}

// Inflate returns the string represented by the given LeagueHeapID
func (id LeagueHeapID) Inflate(db *bolt.DB) string {
	return InflateLeague(id, db)
}

// StashItemModToCompact compacts the given source ItemMod and
// StringHeapID for the mod's text into our internal format
func StashItemModToCompact(mod stash.ItemMod,
	modStringID StringHeapID) ItemMod {

	// Average the values to result in single value
	//
	// 0 Values results in 0 as computed Value
	// 1 Values results in Values[0] as computed Value
	// >1 Values results in Average(Values) as computed Value
	var value uint16
	for _, val := range mod.Values {
		value += val
	}
	// Average
	avgValue := (float64(value) / float64(len(mod.Values)))
	// And scale
	value = uint16(avgValue * ItemModAverageScaleFactor)

	return ItemMod{
		Value: value,
		Mod:   modStringID,
	}

}

// Inflate returns an inflated equivalent item modifier for human use
func (mod ItemMod) Inflate(db *bolt.DB) stash.ItemMod {

	return stash.ItemMod{
		Template: []byte(mod.Mod.Inflate(db)),
		Values:   []uint16{mod.Value},
	}

}

// Inflate returns an inflated equivalent item fit for human use
func (item Item) Inflate(db *bolt.DB) stash.Item {

	// Initialize with the most trivial portions
	fat := stash.Item{
		Name:       item.Name.Inflate(db),
		TypeLine:   item.TypeLine.Inflate(db),
		Note:       item.Note.Inflate(db),
		RootType:   item.RootType.Inflate(db),
		RootFlavor: item.RootFlavor.Inflate(db),
		League:     item.League.Inflate(db),
		Corrupted:  item.Corrupted,
		Identified: item.Identified,
	}

	// And the modifiers
	fatMods := make([]stash.ItemMod, len(item.Mods))
	for i, mod := range item.Mods {
		fatMods[i] = mod.Inflate(db)
	}
	// Uhhhhh.... yeah okay, I don't really care about this.
	// TODO: rethink choices
	fat.ExplicitMods = fatMods

	return fat
}

// StashItemsToCompact converts fat Item records to their compact form
//
// This also ensures all strings present on that item will be available
// on the StringHeap
func StashItemsToCompact(items []stash.Item, when Timestamp,
	db *bolt.DB) ([]Item, error) {

	compact := make([]Item, len(items))

	err := db.Update(func(tx *bolt.Tx) error {
		// Translate leagues on a per-item basis
		leagues := make([]string, len(items))
		for i, item := range items {
			leagues[i] = item.League
		}
		leagueIds, err := setLeagues(leagues, tx)
		if err != nil {
			return errors.Wrap(err, "failed to add leagues to LeagueHeap")
		}

		// Build compact items from the ids and fill in non-StringHeap information
		for i, item := range items {
			compact[i] = Item{
				ID:         ID{}, // Explicitly empty on entrance
				GGGID:      GGGIDFromUID(item.ID),
				Stash:      GGGIDFromUID(item.StashID),
				League:     leagueIds[i],
				Identified: item.Identified,
				Corrupted:  item.Corrupted,
				When:       when,
			}
		}

		// Populate StringHeap related information
		if err := setStringsForItems(items, compact, tx); err != nil {
			return errors.New("failed to set strings")
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to translate item to db form")
	}

	// Then add populate internal IDs, adding them as necessary
	//
	// Left outside the above transaction due to benchmarking results.
	err = GetTranslations(compact, db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add internal IDs to items")
	}

	return compact, nil

}

// StashStashToCompact converts fat Item records to their compact form
// while also stripping items out in their compact form.
func StashStashToCompact(stashes []stash.Stash, when time.Time,
	db *bolt.DB) ([]Stash, [][]Item, error) {

	// Grab a new timestamp, all of the Stashes will share the same time
	whenTS := TimeToTimestamp(when)

	// Compact stashes and flatten items
	compact := make([]Stash, len(stashes))[:0] // Sliced to zero to allow append
	flatItems := make([]stash.Item, 0)
	// Keep track of per-stash leagues
	leagues := make([]string, len(stashes))[:0]
	// Keep track of items per-stash so we can unflatten them into per-stash
	// item sets
	itemsPerStash := make([]int, len(stashes))[:0]
	for _, stash := range stashes {
		// We skip empty stashes as we will be unable to assign
		// then with a LeagueHeapID
		if len(stash.Items) == 0 {
			continue
		}

		compactStash := Stash{
			AccountName: stash.AccountName,
			ID:          GGGIDFromUID(stash.ID),
		}

		// Populate GGGIDs in this Stash
		ids := make([]GGGID, len(stash.Items))
		for i, item := range stash.Items {
			ids[i] = GGGIDFromUID(item.ID)
		}
		compactStash.Items = ids

		compact = append(compact, compactStash)

		// Note the league for this Stash
		leagues = append(leagues, stash.Items[0].League)
		// Note the number of items included in this stash
		itemsPerStash = append(itemsPerStash, len(stash.Items))

		flatItems = append(flatItems, stash.Items...)
	}

	// Fetch and decorate the ids to the compact stashes
	leagueIDs, err := SetLeagues(leagues, db)
	if err != nil {
		err = errors.Wrap(err, "failed to add LeagueHeapIDs to stashes")
		return nil, nil, err
	}
	for i, id := range leagueIDs {
		compact[i].League = id
	}

	// Grab the compact items as their flat form
	compactItems, err := StashItemsToCompact(flatItems, whenTS, db)
	if err != nil {
		err = errors.Wrap(err, "failed to compact items")
		return nil, nil, err
	}

	// Unflatten the items so they can match an associated stash
	items := make([][]Item, len(leagueIDs)) // Sized exactly using leagueIDs
	lastItemBase := 0
	for i, count := range itemsPerStash {
		itemGroup := compactItems[lastItemBase : lastItemBase+count]
		items[i] = itemGroup
		// items[i] = compactItems[lastItemBase:count]
		lastItemBase += count
	}

	return compact, items, nil

}
