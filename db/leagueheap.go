package db

import (
	"fmt"

	"github.com/boltdb/bolt"
)

const leagueHeapBucket string = "leagueHeap"
const leagueHeapInverseBucket string = "leagueHeapInvert"

const leagueNamespaceBucket string = "leagueNamespace"

// All buckets directly under a league are located here.
//
// Any league will always contain these
var leagueSubBuckets = []string{
	itemStoreBucket, indiceBucket, idTranslateBucket, stashBucket,
}

// getLeagueBucket returns the top-level bucket for a specific league
//
// league buckets are stored inside leagueNamespaceBucket as their LeagueHeapID
func getLeagueBucket(league LeagueHeapID, tx *bolt.Tx) *bolt.Bucket {
	// Grab root Bucket
	rootBucket := tx.Bucket([]byte(leagueNamespaceBucket))
	if rootBucket == nil {
		panic(fmt.Sprintf(" %s not found", leagueNamespaceBucket))
	}

	leagueBytes := league.ToBytes()
	leagueBucket := rootBucket.Bucket(leagueBytes)
	if leagueBucket == nil {
		// Add the league bucket itself
		var err error
		leagueBucket, err = rootBucket.CreateBucket(leagueBytes)
		if err != nil {
			panic(fmt.Sprintf("cannot create league bucket, err=%s", err))
		}
		// Add any additional sub-buckets
		if err := addLeagueSubBuckets(leagueBucket, tx); err != nil {
			panic(fmt.Sprintf("cannot create league sub buckets, err=%s", err))
		}
	}
	return leagueBucket
}

// addLeagueSubBuckets adds buckets directly related to a league
//
// This is suger for checkLeague
func addLeagueSubBuckets(leagueBucket *bolt.Bucket, tx *bolt.Tx) error {
	return checkLeague(leagueBucket, tx)
}

// checkLeague ensures that bucekts directly related to a league are present
//
// This should be run before anything else is run in the database as this ensures
// that a league contains registered buckets.
//
// This a an operation that allows league namespaced buckets
// to be registered ahead of time and always exist whenever the league
// bucket itself exists.
func checkLeague(leagueBucket *bolt.Bucket, tx *bolt.Tx) error {
	for _, b := range leagueSubBuckets {
		_, err := leagueBucket.CreateBucketIfNotExists([]byte(b))
		if err != nil {
			return fmt.Errorf("failed to add league bucket, bucket=%s, err=%s", b, err)
		}
	}

	return nil
}

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
	heap.Put([]byte(index), LeagueHeapIDFromSequence(seq).ToBytes())

	// Also add it to the inverseBucket
	var inverter *bolt.Bucket
	if inverter = tx.Bucket([]byte(leagueHeapInverseBucket)); inverter == nil {
		return 0, fmt.Errorf("%s not found", leagueHeapInverseBucket)
	}
	inverter.Put(LeagueHeapIDFromSequence(seq).ToBytes(), []byte(index))

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

// ListLeagues returns a list of all stored leagues
func ListLeagues(db *bolt.DB) ([]string, error) {

	leagues := make([]string, 0)

	return leagues, db.View(func(tx *bolt.Tx) error {

		// Fetch the heap bucket
		var heap *bolt.Bucket
		if heap = tx.Bucket([]byte(leagueHeapBucket)); heap == nil {
			return fmt.Errorf("%s not found", leagueHeapBucket)
		}

		c := heap.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if v != nil {
				leagues = append(leagues, string(k))
			}
		}

		return nil
	})

}

// InflateLeague returns the string represenation of the given LeagueHeapID
//
// If you are providing a valid LeagueHeapID, this should never fail.
func InflateLeague(id LeagueHeapID, db *bolt.DB) string {

	var result string

	db.View(func(tx *bolt.Tx) error {

		// Fetch the inverter bucket
		var inverter *bolt.Bucket
		if inverter = tx.Bucket([]byte(leagueHeapInverseBucket)); inverter == nil {
			panic(fmt.Sprintf("%s does not exist when assumed", leagueHeapInverseBucket))
		}

		// Fetch the string from the inverter
		resultBytes := inverter.Get(id.ToBytes())
		result = string(resultBytes)

		return nil
	})

	return result
}
