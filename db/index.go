package db

import (
	"fmt"

	"bytes"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

const indiceBucket string = "indices"

// getLeagueIndexBucketRO returns the bucket corresponding
// to a specific league's index. This will never write
// and can be used safely with a readonly transaction.
//
// Will either panic or return a valid bucket.
func getLeagueIndexBucket(league LeagueHeapID, tx *bolt.Tx) *bolt.Bucket {
	// Grab league bucket
	leagueBucket := getLeagueBucket(league, tx)

	// This can never fail, its a guarantee that the itemStoreBucket was registered
	// and will always appear on a valid leagueBucket
	indices := leagueBucket.Bucket([]byte(indiceBucket))
	if indices == nil {
		panic(fmt.Sprintf("%s bucket not found when expected", indiceBucket))
	}

	return indices
}

// getItemModIndexBucket returns a bucket which a given mod can be put
// when considering the item containing it.
//
// This WILL write if a bucket is not found. Hence, readonly tx unsafe.
func getItemModIndexBucket(rootType, rootFlavor, mod StringHeapID,
	league LeagueHeapID, tx *bolt.Tx) (*bolt.Bucket, error) {

	var err error

	// Start at the index bucket
	indexBucket := getLeagueIndexBucket(league, tx)
	rootTypeBucket := indexBucket.Bucket(rootType.ToBytes())
	if rootTypeBucket == nil {
		rootTypeBucket, err = indexBucket.CreateBucket(rootType.ToBytes())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to add index rootType bucket")
		}
	}

	rootFlavorBucket := rootTypeBucket.Bucket(rootFlavor.ToBytes())
	if rootFlavorBucket == nil {
		rootFlavorBucket, err = rootTypeBucket.CreateBucket(rootFlavor.ToBytes())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to add index rootFlavor bucket")
		}
	}

	modBucket := rootFlavorBucket.Bucket(mod.ToBytes())
	if modBucket == nil {
		modBucket, err = rootFlavorBucket.CreateBucket(mod.ToBytes())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to add index mod bucket")
		}
	}

	// If we made it through, our currentBucket should be the one we want
	return modBucket, nil
}

// getItemModIndexBucketRO returns a bucket which a given mod can be put
// when considering the item containing it.
//
// This WILL NOT write if a bucket is not found. Hence, readonly tx unsafe.
func getItemModIndexBucketRO(rootType, rootFlavor, mod StringHeapID,
	league LeagueHeapID, tx *bolt.Tx) (*bolt.Bucket, error) {

	// Start at the index bucket
	indexBucket := getLeagueIndexBucket(league, tx)

	rootTypeBucket := indexBucket.Bucket(rootType.ToBytes())
	if rootTypeBucket == nil {
		return nil, errors.New("invalid rootType, no bucket found")
	}

	rootFlavorBucket := rootTypeBucket.Bucket(rootFlavor.ToBytes())
	if rootFlavorBucket == nil {
		return nil, errors.New("invalid rootFlavor, no bucket found")
	}

	modBucket := rootFlavorBucket.Bucket(mod.ToBytes())
	if modBucket == nil {
		return nil, errors.New("invalid mod, no bucket found")
	}

	// If we made it through, our currentBucket should be the one we want
	return modBucket, nil
}

// ModIndexKeySuffixLength allows us to fetch variable numbers
// of pre-pended values given their length.
const ModIndexKeySuffixLength = TimestampSize

// encodeModIndexKey generates a mod key based off of the provided data
//
// The mod index key is generated as [mod.Values..., now, updateSequence]
func encodeModIndexKey(mod ItemMod, now Timestamp) []byte {

	// Pre-allocate index key so the entire key can be
	// encoded with a single allocation.
	modsLength := 2
	indexKey := make([]byte, ModIndexKeySuffixLength+modsLength)

	// Generate the suffix
	suffix := (indexKey[modsLength:])[:0] // Deal with pre-allocated space
	suffix = append(suffix, now.TruncateToIndexBucket()[:]...)

	if len(suffix) != ModIndexKeySuffixLength {
		panic(fmt.Sprintf("unexpected suffix length, got %d, expected %d",
			len(suffix), ModIndexKeySuffixLength))
	}

	// Fill in the index from the front
	//
	// TODO: avoid appends, pre-size the backing slice to accomodate the
	// contents including the header
	index := indexKey[:0] // Deal with pre-allocated space
	index = append(index, i16tob(mod.Value)...)

	// And return the index with its suffix
	return append(index, suffix...)
}

// decodeModIndexKey decodes a provided mod index key
//
// This returns the values encoded in the key.
//
// This is possible as the suffix is a fixed length and format while
// the values of the modifer are simple appended
func decodeModIndexKey(key []byte) ([]uint16, error) {

	// Basic sanity check
	if len(key) < ModIndexKeySuffixLength {
		return nil, errors.New("invalid index key passed, less than length of suffix")
	}

	// Ensure we are divisible by 2 following the removal of the suffix
	if (len(key)-ModIndexKeySuffixLength)%2 != 0 {
		return nil, errors.New("invalid index key passed, values malformed")
	}

	valueBytes := key[:len(key)-ModIndexKeySuffixLength]
	values := make([]uint16, len(valueBytes)/2)
	for index := 0; index*2 < len(valueBytes); index++ {
		values[index] = btoi16(valueBytes[index*2:])
	}

	return values, nil

}

// IndexItems adds tbe given items to their correct indices
// for efficient lookup. Returns number of index entries added.
//
// Provided items CAN differ in their league.
func IndexItems(items []Item, tx *bolt.Tx) (int, error) {

	// Sanity check passed in transaction, better to do this than panic.
	if !tx.Writable() {
		return 0, errors.New("cannot IndexItems on readonly transaction")
	}

	// Silently exit when no items present to add
	if len(items) < 1 {
		return 0, nil
	}

	var added int

	for _, item := range items {

		for _, mod := range item.Mods {
			// Grab the bucket we can actually insert things into

			itemModBucket, err := getItemModIndexBucket(item.RootType, item.RootFlavor,
				mod.Mod, item.League, tx)
			if err != nil {
				return 0, errors.New("failed to get item mod bucket")
			}

			modKey := encodeModIndexKey(mod, item.When)

			// Check for pre-existing items in the bucket, if none, we establish
			// the bucket
			existing := itemModBucket.Get(modKey)
			wrapped := IndexEntry(existing)
			wrapped = IndexEntryAppend(wrapped, item.ID)
			itemModBucket.Put(modKey, wrapped)
			added++
		}
	}

	return added, nil

}

// DeindexItems removes tbe given items from their correct indices
//
// If an index entry cannot be removed, we return an error. This ensures
// all existing index entries must be alive
func DeindexItems(items []Item, tx *bolt.Tx) error {

	// Sanity check passed in transaction, better to do this than panic.
	if !tx.Writable() {
		return errors.New("cannot IndexItems on readonly transaction")
	}

	// Silently exit when no items present to add
	if len(items) < 1 {
		return nil
	}

	for _, item := range items {

		for _, mod := range item.Mods {
			// Grab the bucket we can actually insert things into
			itemModBucket, err := getItemModIndexBucket(item.RootType, item.RootFlavor,
				mod.Mod, item.League, tx)
			if err != nil {
				return errors.New("failed to get item mod bucket")
			}

			modKey := encodeModIndexKey(mod, item.When)
			existing := itemModBucket.Get(modKey)
			wrapped := IndexEntry(existing)
			wrapped = IndexEntryRemove(wrapped, item.ID)
			if wrapped == nil {
				// Nothing else resides at this index
				itemModBucket.Delete(modKey)
			} else {
				// Add back with removed data
				itemModBucket.Put(modKey, wrapped)
			}
		}
	}

	return nil

}

// IndexEntry represents bytes interpreted as an entry within the index
//
// Whenever possible, we avoid allocations.
type IndexEntry []byte

// IndexEntryAppend adds another ID to the entry
//
// If an id is already present in the id, we end up with a duplicate.
// Such is life.
func IndexEntryAppend(entry IndexEntry, id ID) IndexEntry {
	var result []byte
	if entry == nil {
		// Copy necessary due to boltdb semantics for passed buffers
		result = make([]byte, len(id))
		copy(result, id[:])
	} else {
		// We assume item not already present in bucket.
		// If it is, we end up with a duplicate.
		//
		// Allocate a buffer large enough for an append
		// without another allocation.
		// Yes, this looks super dirty. TODO: cleanup D:
		result = make([]byte, len(entry)+IDSize)[:0]
		result = append(result, entry...)
		result = append(result, id[:]...)
	}

	return IndexEntry(result)
}

// IndexEntryRemove removes a given ID from the entry
//
// If the ID is the last of the entry, the backing slice is set
// to nil. In that case, its the callers responsibility to ensure they
// Unwrap a valid byte slice.
func IndexEntryRemove(entry IndexEntry, id ID) IndexEntry {

	// If the backing array is nil, then we can't remove an ID
	// and the database is inconsistent.
	if entry == nil {
		panic(fmt.Sprintf("attempted to remove ID from nil IndexEntry, id=%v",
			id))
	}

	// For removal, we Stride over the array in IDSize increments
	// and call Equal to determine which index we remove.
	index := -1 // Index is in terms of IDSize increments
	for i := 0; i < len(entry); i += IDSize {
		equal := bytes.Equal(id[:], entry[i:i+IDSize])
		if equal {
			index = i
			break
		}
	}

	// If we can't find the ID, invalid index state, so panic.
	if index == -1 {
		panic(fmt.Sprintf("attempted to remove non-existent ID, id=%v", id))
	}

	// Check if this is the last entry, if yes, then easy nil.
	if len(entry) == IDSize {
		return nil
	}
	// Remove the entry using fewest allocations possible.
	//
	// We have to asssume our internal buffer for the entry came from
	// bolt, hence the new buffer to mutate.
	removed := make([]byte, len(entry)-IDSize)[:0]
	removed = append(removed, entry[:index]...)
	removed = append(removed, entry[index+IDSize:]...)
	return removed
}

// ForEachID calls the provided callback with each id contained
// within the IndexEntry. Order of returned IDs is in Append-descending
func (entry IndexEntry) ForEachID(cb func(id ID)) {
	var id ID
	for i := 0; i < len(entry); i += IDSize {
		copy(id[:], entry[i:i+IDSize])
		cb(id)
	}
}

// GetIDs returns all IDs in the entry.
//
// Provided array slice will be resized if necessary or a new one
// will be created if passed nil. Updated slice will be returned.
func (entry IndexEntry) GetIDs(ids []ID) []ID {

	idCount := len(entry) / IDSize
	if ids == nil || cap(ids) < idCount {
		ids = make([]ID, idCount)
	}
	ids = ids[:0]

	for i := 0; i < len(entry); i += IDSize {
		var id ID
		copy(id[:], entry[i:i+IDSize])
		ids = append(ids, id)
	}

	return ids
}

// IndexEntryCount returns the number of index entries across all leagues
func IndexEntryCount(db *bolt.DB) (int, error) {
	var count int

	leagueStrings, err := ListLeagues(db)
	if err != nil {
		return 0, err
	}
	leagueIDs, err := GetLeagues(leagueStrings, db)
	if err != nil {
		return 0, err
	}

	return count, db.View(func(tx *bolt.Tx) error {

		for _, id := range leagueIDs {
			b := getLeagueIndexBucket(id, tx)
			if b == nil {
				return errors.Errorf("%s bucket not found", itemStoreBucket)
			}
			stats := b.Stats()
			count += stats.KeyN
		}

		return nil
	})
}
