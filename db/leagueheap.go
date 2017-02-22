package db

import (
	"fmt"

	"github.com/boltdb/bolt"
)

const leagueHeapBucket string = "leagueHeap"
const leagueHeapInverseBucket string = "leagueHeapInvert"

// Set a league value in the heap and returns its corresponding LeagueHeapID
//
// A transaction is passed in to allow batch entry
func setLeague(index string, tx *bolt.Tx) (LeagueHeapID, error) {

	// Fetch the heap bucket
	var heap *bolt.Bucket
	if heap = tx.Bucket([]byte(leagueHeapBucket)); heap == nil {
		return 0, fmt.Errorf("%s not found", leagueHeapBucket)
	}

	// If it already exists, early exit
	if result := heap.Get([]byte(index)); result != nil {
		return LeagueHeapIDFromBytes(result), nil
	}

	// If it doesn't, we need a sequence number
	seq, err := heap.NextSequence()
	if err != nil {
		return 0, fmt.Errorf("failed to get NextSequence in %s", leagueHeapBucket)
	}
	heap.Put([]byte(index), i16tob(uint16(seq)))

	// Also add it to the inverseBucket
	var inverter *bolt.Bucket
	if inverter = tx.Bucket([]byte(leagueHeapInverseBucket)); inverter == nil {
		return 0, fmt.Errorf("%s not found", leagueHeapInverseBucket)
	}
	inverter.Put(i16tob(uint16(seq)), []byte(index))

	return LeagueHeapID(seq), nil

}

func getLeague(index string, tx *bolt.Tx) (LeagueHeapID, error) {
	// Fetch the heap bucket
	var heap *bolt.Bucket
	if heap = tx.Bucket([]byte(leagueHeapBucket)); heap == nil {
		return 0, fmt.Errorf("%s not found", leagueHeapBucket)
	}

	// If it already exists, early exit
	var result []byte
	if result = heap.Get([]byte(index)); result == nil {
		return 0, fmt.Errorf("failed to find index in league heap")
	}

	return LeagueHeapIDFromBytes(result), nil
}

// SetLeagues fills in any missing league heap index values
// and maps all indice league values to their corresponding LeagueHeapID
func SetLeagues(leagues []string, db *bolt.DB) ([]LeagueHeapID, error) {
	ids := make([]LeagueHeapID, len(leagues))

	return ids, db.Update(func(tx *bolt.Tx) error {

		for i, index := range leagues {
			id, err := setLeague(index, tx)
			if err != nil {
				return err
			}
			ids[i] = id
		}

		return nil
	})
}

// GetLeagues maps all indice league values to their corresponding LeagueHeapID
func GetLeagues(leagues []string, db *bolt.DB) ([]LeagueHeapID, error) {
	ids := make([]LeagueHeapID, len(leagues))

	return ids, db.View(func(tx *bolt.Tx) error {

		for i, index := range leagues {
			id, err := getLeague(index, tx)
			if err != nil {
				return err
			}
			ids[i] = id
		}

		return nil
	})
}
