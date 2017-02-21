package db

import (
	"encoding/binary"
	"fmt"

	"github.com/Everlag/gothing/stash"
	"github.com/boltdb/bolt"
)

// DBLocation is the on-disk file containing our database
const DBLocation string = "poe.db"

const propertyNameBucket string = "propertyNames"

var bucketNames = [...]string{
	propertyNameBucket,
	stringHeapBucket, stringHeapInverseBucket,
}

// itob returns an 8-byte big endian representation of v.
// Courtesy of boltdb dev logs
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// btoi returns a PutUint64 from its 8-byte big endian representation.
// Courtesy of boltdb dev logs
func btoi(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
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

// addPropertyName adds a property to the bucket
func addPropertyName(property string, tx *bolt.Tx) error {

	b := tx.Bucket([]byte(propertyNameBucket))
	if b == nil {
		return fmt.Errorf("propertyNameBucket bucket not found")
	}

	index := b.Get([]byte(property))
	if index != nil {
		return nil
	}

	seq, err := b.NextSequence()
	if err != nil {
		return fmt.Errorf("failed to get NextSequence in propertyNameBucket")
	}
	b.Put([]byte(property), itob(seq))

	return nil
}

// AddPropertyNamesFromResponse adds all property names found in a stash
// response to the database
func AddPropertyNamesFromResponse(resp *stash.Response, db *bolt.DB) error {
	return db.Batch(func(tx *bolt.Tx) error {
		// Such iteration D:
		for _, stash := range resp.Stashes {
			for _, item := range stash.Items {
				for _, property := range item.Properties {
					if err := addPropertyName(property.Name, tx); err != nil {
						return fmt.Errorf("failed to add property name, err=%s", err)
					}
				}
			}
		}
		return nil
	})
}

// GetPropertyID returns the integer value associated with a property name
func GetPropertyID(property string, db *bolt.DB) (uint64, error) {
	var index uint64

	return index, db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(propertyNameBucket))
		if b == nil {
			return fmt.Errorf("propertyNameBucket bucket not found")
		}

		indexBytes := b.Get([]byte(property))
		if indexBytes == nil {
			return fmt.Errorf("property not found")
		}

		index = btoi(indexBytes)

		return nil
	})
}

// PropertyNameCount returns the number of defined properties
func PropertyNameCount(db *bolt.DB) (int, error) {
	var count int

	return count, db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(propertyNameBucket))
		if b == nil {
			return fmt.Errorf("propertyNameBucket bucket not found")
		}

		stats := b.Stats()
		count = stats.KeyN

		return nil
	})
}

// Boot gets the database from disk and performs necessary setup
func Boot() (*bolt.DB, error) {
	db, err := bolt.Open(DBLocation, 777, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s as boltdb, err=%s", DBLocation, err)
	}

	if err := setupBuckets(db); err != nil {
		return nil, fmt.Errorf("failed to setup buckets, err=%s", err)
	}
	return db, nil
}
