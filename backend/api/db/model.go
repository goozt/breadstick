package db

import (
	"github.com/gofiber/storage/bbolt"
	"github.com/gofiber/utils"
	bolt "go.etcd.io/bbolt"
)

type DB struct {
	db   *bbolt.Storage
	conn *bolt.DB
}

// Create new database instance
func NewDB(path, initialBucket string) *DB {
	db := bbolt.New(bbolt.Config{
		Database: path,
		Bucket:   initialBucket,
		Reset:    false,
	})
	return &DB{db: db, conn: db.Conn()}
}

// Create new store in database
func (db *DB) NewStore(name string) error {
	return db.conn.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(utils.UnsafeBytes(name))
		return err
	})
}

// Get store instance by name
func (db *DB) GetStore(name string) *Store {
	return &Store{name, db.conn}
}

// Close database connection
func (db *DB) Close() error {
	return db.db.Close()
}
