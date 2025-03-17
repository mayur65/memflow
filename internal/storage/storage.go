package storage

import (
	"fmt"
	"github.com/mayur65/memflow/internal/config"
	"sync"
	"time"
)

type DB struct {
	data map[string]string
	ttl  map[string]time.Time
	mu   sync.RWMutex
}

func InitDB() *DB {
	return &DB{
		data: make(map[string]string),
		ttl:  make(map[string]time.Time),
	}
}

func (db *DB) Get(key string) string {
	db.mu.RLock()
	defer db.mu.RUnlock()

	expiry, ok := db.ttl[key]

	if !ok {
		return "KEY_NOT_FOUND"
	}

	if time.Now().After(expiry) {
		return "Key expired"
	}

	val, ok := db.data[key]

	return val
}

func (db *DB) Set(key, value string) string {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[key] = value
	db.ttl[key] = time.Now().Add(config.TimeToLive)

	fmt.Println(db.data)

	return "200 - Value set for Key"
}

func (db *DB) Delete(key string) string {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, ok := db.data[key]

	if !ok {
		return "KEY_NOT_FOUND"
	}

	delete(db.data, key)
	delete(db.ttl, key)

	return "200 - Value deleted for Key"
}

func (db *DB) PeriodicCleaning() {

	for {
		time.Sleep(time.Minute)

		db.mu.Lock()

		now := time.Now()

		for k, v := range db.ttl {
			if now.After(v) {
				delete(db.data, k)
				delete(db.ttl, k)
			}
		}

		db.mu.Unlock()
	}

}
