package db

import (
	"fmt"

	"time"

	"github.com/boltdb/bolt"
)

const updateSnapshotHistoryBuckets = "updateSnapshotHistory"

// RecordChangeIDWasEntered adds a given changeID to the database
// with its value being the json time it was entered
func RecordChangeIDWasEntered(changeID string, db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(updateSnapshotHistoryBuckets))
		if b == nil {
			return fmt.Errorf("propertyNameBucket bucket not found")
		}

		// Ensure this was not previously entered
		if value := b.Get([]byte(changeID)); value != nil {
			return fmt.Errorf("previous entry time exists for changeID")
		}

		nowJSON, err := time.Now().MarshalBinary()
		if err != nil {
			return fmt.Errorf("failed to marshal time.Now to json")
		}
		b.Put([]byte(changeID), nowJSON)

		return nil
	})
}

// UpdateWasEntered determines if a given update has been entered into the database
func UpdateWasEntered(changeID string, db *bolt.DB) (bool, error) {
	entered := false

	return entered, db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(updateSnapshotHistoryBuckets))
		if b == nil {
			return fmt.Errorf("propertyNameBucket bucket not found")
		}

		value := b.Get([]byte(changeID))
		entered = value != nil
		return nil
	})
}
