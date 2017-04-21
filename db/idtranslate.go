package db

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

const idTranslateBucket string = "idTranslator"

// getIDTranslateItemBucket returns the bucket corresponding
// to a specific league for translating GGG identifiers to ours
//
// Will either panic or return a valid bucket.
func getIDTranslateItemBucket(league LeagueHeapID, tx *bolt.Tx) *bolt.Bucket {
	// Grab league bucket
	leagueBucket := getLeagueBucket(league, tx)

	// Grab the translator
	//
	// This can never fail, its a guarantee that the idTranslateBucket was registered
	// and will always appear on a valid leagueBucket
	translator := leagueBucket.Bucket([]byte(idTranslateBucket))
	if translator == nil {
		panic(fmt.Sprintf("%s bucket not found when expected", idTranslateBucket))
	}

	return translator
}

// getTranslation returns the internal ID for a provided GGG identifier
// and will allocate a new identifier if none are found.
func getTranslation(league LeagueHeapID,
	external GGGID, tx *bolt.Tx) (ID, error) {
	// Fetch the heap bucket
	var translator *bolt.Bucket

	if translator = getIDTranslateItemBucket(league, tx); translator == nil {
		return ID{},
			errors.Errorf("translator bucket for league not found, league=%d",
				league)
	}

	// If it already exists, early exit
	if result := translator.Get(external[:]); result != nil {
		var id ID
		copy(id[:], result)
		return id, nil
	}

	// Assign it a new sequence number as necessary
	seq, err := translator.NextSequence()
	if err != nil {
		return ID{}, errors.Errorf("failed to get NextSequence in %s",
			idTranslateBucket)
	}
	id := IDFromSequence(seq)
	translator.Put(external[:], id[:])

	return id, nil
}

// GetTranslations associates each provided item with an interal ID
// if it has not already been assigned one.
//
// This modifies the provided items if they are assigned an ID
func GetTranslations(items []Item, db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {

		for i, item := range items {
			id, err := getTranslation(item.League, item.GGGID, tx)
			if err != nil {
				return err
			}

			item.ID = id
			items[i] = item
		}

		return nil
	})
}

// GetGGGIDTranslations associates each provided item with an interal ID
// if it has not already been assigned one.
//
// All IDs are in the context of the provided league, hence this is typically
// used for managing individual stash updates and that's why
// we take a transaction.
//
// This returns the newly translated ID positionally mapping to the provided
// GGGIDs.
func GetGGGIDTranslations(gggs []GGGID, league LeagueHeapID,
	tx *bolt.Tx) ([]ID, error) {

	ids := make([]ID, len(gggs))

	for i, ggg := range gggs {
		id, err := getTranslation(league, ggg, tx)
		if err != nil {
			return nil, err
		}

		ids[i] = id
	}

	return ids, nil
}
