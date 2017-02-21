package db

import (
	"fmt"

	"github.com/boltdb/bolt"
)

const stringHeapBucket string = "stringHeap"
const stringHeapInverseBucket string = "stringHeapInvert"

// Set a string value in a heap and returns its corresponding StringHeapID
//
// A transaction is passed in to allow batch entry
func setString(index string, tx *bolt.Tx) (StringHeapID, error) {

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
	seq, err := heap.NextSequence()
	if err != nil {
		return 0, fmt.Errorf("failed to get NextSequence in %s", stringHeapBucket)
	}
	heap.Put([]byte(index), itob(seq))

	// Also add it to the inverseBucket
	var inverter *bolt.Bucket
	if inverter = tx.Bucket([]byte(stringHeapInverseBucket)); inverter == nil {
		return 0, fmt.Errorf("%s not found", stringHeapInverseBucket)
	}
	inverter.Put(itob(seq), []byte(index))

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
