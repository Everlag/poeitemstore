package db

import (
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

// ItemStoreQuery represents a query running naively over the ItemStore
//
// This is a naive and inefficient query. This is mostly used for ensuring
// the correctness of more efficient query methods.
//
// An ItemStoreQuery can be rerun by reinitializing the ctx; this
// typically happens when the query is Run.
type ItemStoreQuery struct {
	// Type and flavor of the item we're looking up
	rootType, rootFlavor StringHeapID
	// Minimum mod values we are required to find
	// are pointed to by their StringHeapID for easy lookup
	minModMap map[StringHeapID]uint16
	// League we are searching for
	league LeagueHeapID
	// How many items we are limited to finding
	maxDesired int
}

// NewItemStoreQuery returns an ItemStoreQuery with no context
//
// If len(mods) != len(minModValues), we panic; so don't give us garbage
func NewItemStoreQuery(rootType, rootFlavor StringHeapID,
	mods []StringHeapID, minModValues []uint16,
	league LeagueHeapID,
	maxDesired int) ItemStoreQuery {

	minModMap := make(map[StringHeapID]uint16)
	for i, mod := range mods {
		minModMap[mod] = minModValues[i]
	}

	return ItemStoreQuery{
		rootType, rootFlavor,
		minModMap,
		league, maxDesired,
	}

}

// checkItem determines if a given item satisfies the query
func (q *ItemStoreQuery) checkItem(item Item) bool {

	// Perform trivial check before expensive mod check
	validRoot := q.rootType == item.RootType
	validFlavor := q.rootFlavor == item.RootFlavor
	if !(validRoot && validFlavor) {
		return false
	}

	// Check each mod present on the provided item
	// against the mods we need.
	countPresent := 0
	for _, mod := range item.Mods {
		required, ok := q.minModMap[mod.Mod]
		if !ok {
			continue
		}
		if len(mod.Values) < 1 || mod.Values[0] > required {
			countPresent++
		}
	}

	return countPresent >= len(q.minModMap)
}

// checkPair determines if a pair is acceptable for our query
func (q *ItemStoreQuery) checkPair(k, v []byte) (bool, error) {
	var item Item
	_, err := item.UnmarshalMsg(v)
	if err != nil {
		return false, errors.Wrap(err, "failed to UnmarshalMsg itemstore item")
	}
	return q.checkItem(item), nil
}

// Run initialises transaction context for a query and attempts
// to find desired items.
func (q *ItemStoreQuery) Run(db *bolt.DB) ([]ID, error) {

	// Preallocate the ids to fit the max we want but also
	// allow us to use append rather than deal with indices
	ids := make([]ID, q.maxDesired)[:0]

	err := db.View(func(tx *bolt.Tx) error {

		b := getLeagueItemBucket(q.league, tx)
		if b == nil {
			return errors.Errorf("failed to get league item bucket, LeagueHeapID=%d",
				q.league)
		}

		// Grab and set the cursor to last
		c := b.Cursor()
		k, v := c.Last()
		if k == nil {
			return errors.New("failed to get last item in itemstore, empty bucket")
		}
		// Test the item we got back as long as it isn't a bucket
		if len(v) > 0 {
			valid, err := q.checkPair(k, v)
			if err != nil {
				return errors.Wrap(err, "failed to check pair in itemstore")
			}
			if valid {
				var id ID
				copy(id[:], k)
				ids = append(ids, id)
			}
		}

		// Perform the actual search along the itemstore
		//
		// We go until we exhaust the entire store or find as many as we need
		for index := 0; len(ids) < q.maxDesired; index++ {

			// Grab a pair
			k, v := c.Prev()
			// Ignore nested buckets but also
			// handle reaching the start of the bucket
			if len(v) == 0 {
				// Both nil means we're done
				if k == nil {
					break
				}
				continue
			}
			valid, err := q.checkPair(k, v)
			if err != nil {
				return errors.Wrap(err, "failed to check pair in itemstore")
			}
			if valid {
				var id ID
				copy(id[:], k)
				ids = append(ids, id)
			}
		}

		return nil
	})

	return ids, err
}
