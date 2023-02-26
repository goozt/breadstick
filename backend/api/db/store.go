package db

import (
	"github.com/gofiber/utils"
	"go.etcd.io/bbolt"
)

type Store struct {
	Name string
	conn *bbolt.DB
}

// Delete stored instance
func (s *Store) DeleteStore() error {
	return s.conn.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket(utils.UnsafeBytes(s.Name))
	})
}

// Reset stores instance
func (s *Store) ResetStore() error {
	err := s.DeleteStore()
	if err != nil {
		return err
	}
	return s.conn.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(utils.UnsafeBytes(s.Name))
		return err
	})
}

// Get internal bucket instance by name
func (s *Store) getBucket(tx *bbolt.Tx) *bbolt.Bucket {
	return tx.Bucket(utils.UnsafeBytes(s.Name))
}

// Get value by key from store instance
func (s *Store) Get(key string) ([]byte, error) {
	var value []byte
	err := s.conn.View(func(tx *bbolt.Tx) error {
		b := s.getBucket(tx)
		value = b.Get(utils.UnsafeBytes(key))
		return nil
	})
	return value, err
}

// Get all key-value pairs from store
func (s *Store) GetAll() (map[string][]byte, error) {
	values := make(map[string][]byte)
	err := s.conn.View(func(tx *bbolt.Tx) error {
		b := s.getBucket(tx)

		b.ForEach(func(k, v []byte) error {
			values[string(k)] = v
			return nil
		})
		return nil
	})
	return values, err
}

// Set value by key to store instance
func (s *Store) Set(key string, value []byte) error {
	if len(key) <= 0 || len(value) <= 0 {
		return nil
	}

	return s.conn.Update(func(tx *bbolt.Tx) error {
		b := s.getBucket(tx)
		return b.Put(utils.UnsafeBytes(key), value)
	})
}

// Check if given key exists
func (s *Store) KeyExists(key string) bool {
	v, e := s.Get(key)
	if e != nil || v == nil {
		return false
	}
	return true
}

// Check if given hash exists
func (s *Store) ValueExists(value string, valueFn func([]byte) (string, error)) bool {
	m, e := s.GetAll()
	if e != nil || m == nil || len(m) < 1 {
		return false
	}

	for _, v := range m {
		itemValue, err := valueFn(v)
		if err != nil {
			continue
		}
		if itemValue == value {
			return true
		}
	}

	return false
}

// Delete value by key from store instance
func (s *Store) Delete(key string) error {
	return s.conn.Update(func(tx *bbolt.Tx) error {
		b := s.getBucket(tx)
		return b.Delete(utils.UnsafeBytes(key))
	})
}
