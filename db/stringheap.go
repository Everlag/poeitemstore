package db

import (
	"fmt"

	"github.com/Everlag/poeitemstore/stash"
	"github.com/boltdb/bolt"
)

const stringHeapBucket string = "stringHeap"
const stringHeapInverseBucket string = "stringHeapInvert"

// Set a string value in a heap and returns its corresponding StringHeapID
//
// A transaction is passed in to allow batch entry
func setString(index string, tx *bolt.Tx) (StringHeapID, error) {

	// Sanity check index, we should never receive an empty value
	if len(index) == 0 {
		return 0, fmt.Errorf("zero length index passed to setString")
	}

	// Fetch the heap bucket
	var heap *bolt.Bucket
	if heap = tx.Bucket([]byte(stringHeapBucket)); heap == nil {
		return 0, fmt.Errorf("%s not found", stringHeapBucket)
	}

	// If it already exists, early exit
	if result := heap.Get([]byte(index)); result != nil {
		return StringHeapIDFromBytes(result), nil
	}

	// If it doesn't, we need a sequence number
	// fmt.Println("need sequence number for unseen string", index, len(index))
	seq, err := heap.NextSequence()
	if err != nil {
		return 0, fmt.Errorf("failed to get NextSequence in %s", stringHeapBucket)
	}
	heap.Put([]byte(index), i32tob(uint32(seq)))

	// Also add it to the inverseBucket
	var inverter *bolt.Bucket
	if inverter = tx.Bucket([]byte(stringHeapInverseBucket)); inverter == nil {
		return 0, fmt.Errorf("%s not found", stringHeapInverseBucket)
	}
	inverter.Put(i32tob(uint32(seq)), []byte(index))

	return StringHeapID(seq), nil

}

func getString(index string, tx *bolt.Tx) (StringHeapID, error) {
	// Fetch the heap bucket
	var heap *bolt.Bucket
	if heap = tx.Bucket([]byte(stringHeapBucket)); heap == nil {
		return 0, fmt.Errorf("%s not found", stringHeapBucket)
	}

	// If it already exists, early exit
	var result []byte
	if result = heap.Get([]byte(index)); result == nil {
		return 0, fmt.Errorf("failed to find index in string heap")
	}

	return StringHeapIDFromBytes(result), nil
}

// SetStrings fills in any missing string heap index values
// and maps all indice string values to their corresponding StringHeapID
func SetStrings(indices []string, db *bolt.DB) ([]StringHeapID, error) {
	ids := make([]StringHeapID, len(indices))

	return ids, db.Update(func(tx *bolt.Tx) error {

		for i, index := range indices {
			id, err := setString(index, tx)
			if err != nil {
				return err
			}
			ids[i] = id
		}

		return nil
	})
}

// SetStringsForItems fills in any missing string heap index values
// and maps all strings present on provided items to their
// StringHeapID.
//
// This will call the gen function until it returns no further non-nil values.
// The gen function is expected to return the Items to add and pre-initialized
// Items to decorate with strings
//
// This function's use case is to allow many, separate sets of strings
// to be added while retaining their index positions.
func SetStringsForItems(source []stash.Item, target []Item, db *bolt.DB) error {

	if len(source) != len(target) {
		return fmt.Errorf("length of provided source does not match target, %d!=%d",
			len(source), len(target))
	}
	return db.Update(func(tx *bolt.Tx) error {

		var err error
		for i, item := range source {
			// Repetitive and easy
			target[i].Name, err = setString(item.Name, tx)
			if err != nil {
				return err
			}
			target[i].TypeLine, err = setString(item.TypeLine, tx)
			if err != nil {
				return err
			}
			target[i].Note, err = setString(item.Note, tx)
			if err != nil {
				return err
			}
			target[i].RootFlavor, err = setString(item.RootFlavor, tx)
			if err != nil {
				return err
			}
			target[i].RootType, err = setString(item.RootType, tx)
			if err != nil {
				return err
			}

			// Fill in mods, nasty but fast
			sourceMods := item.GetMods()
			target[i].Mods = make([]ItemMod, len(sourceMods))
			for k, mod := range sourceMods {
				target[i].Mods[k].Mod, err = setString(string(mod.Template),
					tx)
				if err != nil {
					return err
				}
				target[i].Mods[k].Values = mod.Values
			}
		}

		return nil
	})
}

// GetStrings maps all indice string values to their corresponding StringHeapID
func GetStrings(indices []string, db *bolt.DB) ([]StringHeapID, error) {
	ids := make([]StringHeapID, len(indices))

	return ids, db.View(func(tx *bolt.Tx) error {

		for i, index := range indices {
			id, err := getString(index, tx)
			if err != nil {
				return err
			}
			ids[i] = id
		}

		return nil
	})
}

// InflateString returns the string represenation of the given StringHeapID
//
// If you are providing a valid StringHeapID, this should never fail.
func InflateString(id StringHeapID, db *bolt.DB) string {

	var result string

	db.View(func(tx *bolt.Tx) error {

		// Fetch the inverter bucket
		var inverter *bolt.Bucket
		if inverter = tx.Bucket([]byte(stringHeapInverseBucket)); inverter == nil {
			panic(fmt.Sprintf("%s does not exist when assumed", stringHeapInverseBucket))
		}

		// Fetch the string from the inverter
		resultBytes := inverter.Get(id.ToBytes())
		result = string(resultBytes)

		return nil
	})

	return result
}
