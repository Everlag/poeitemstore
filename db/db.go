package db

import (
	"encoding/binary"
	"fmt"

	"github.com/boltdb/bolt"
)

// DBLocation is the on-disk file containing our database
const DBLocation string = "poe.db"

var bucketNames = [...]string{
	stringHeapBucket, stringHeapInverseBucket,
	leagueHeapBucket, leagueHeapInverseBucket,
	updateSnapshotHistoryBuckets,
	leagueNamespaceBucket,
}

// i64tob returns an 8-byte big endian representation of v.
// Courtesy of boltdb dev logs
func i64tob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// btoi64 returns a PutUint64 from its 8-byte big endian representation.
// Courtesy of boltdb dev logs
func btoi64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// i32tob returns a 4-byte big endian representation of v.
// Courtesy of boltdb dev logs
func i32tob(v uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	return b
}

// btoi32 returns a PutUint32 from its 4-byte big endian representation.
// Courtesy of boltdb dev logs
func btoi32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

// i16tob returns an 2-byte big endian representation of v.
// Courtesy of boltdb dev logs
func i16tob(v uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, v)
	return b
}

// btoi returns a PutUint16 from its 2-byte big endian representation.
// Courtesy of boltdb dev logs
func btoi16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

func setupBuckets(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		for _, bucket := range bucketNames {
			_, err := tx.CreateBucketIfNotExists([]byte(bucket))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
		}
		return nil
	})
}

// Boot gets the database from disk and performs necessary setup
//
// If path is empty, it uses the default DBLocation
func Boot(path string) (*bolt.DB, error) {
	if path == "" {
		path = DBLocation
	}

	db, err := bolt.Open(path, 777, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s as boltdb, err=%s", DBLocation, err)
	}

	// Ensure root level buckets exist
	if err := setupBuckets(db); err != nil {
		return nil, fmt.Errorf("failed to setup buckets, err=%s", err)
	}

	// Ensure league level buckets exist on each league
	leagueStrings, err := ListLeagues(db)
	if err != nil {
		return nil, fmt.Errorf("failed to list leagues, err=%s", err)
	}
	leagueIDs, err := GetLeagues(leagueStrings, db)
	if err != nil {
		return nil, fmt.Errorf("failed to convert league strings to ids, err=%s", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {

		for _, id := range leagueIDs {
			b := getLeagueItemBucket(id, tx)
			if err = checkLeague(b, tx); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to checkLeagues, err=%s", err)
	}

	return db, nil
}
