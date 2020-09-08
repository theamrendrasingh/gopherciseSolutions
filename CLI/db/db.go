package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var bucketName = []byte("tasks")
var db *bolt.DB

type TaskItem struct {
	Id    int
	Value string
}

func OpenDB(path string) error {
	var err error
	db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(bucketName)
		return err
	})
}

func AddTask(task string) (int, error) {
	var idx int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		id, _ := b.NextSequence()
		idx = int(id)
		key := intToByte(int(id))
		return b.Put(key, []byte(task))
	})

	if err != nil {
		return -1, err
	}
	return idx, nil
}

func DeleteItem(id int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)

		return b.Delete(intToByte(id))
	})

	return err
}

func ListTaskItems() ([]TaskItem, error) {
	var items []TaskItem

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			items = append(items, TaskItem{
				Id:    byteToInt(k),
				Value: string(v),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return items, nil
}

func intToByte(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func byteToInt(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
