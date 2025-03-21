package storage

import (
	"encoding/gob"
	"github.com/mayur65/memflow/internal/config"
	"log"
	"os"
	"sync"
	"time"
)

type DB struct {
	Data map[string]string
	Ttl  map[string]time.Time
	mu   sync.RWMutex
}

func InitDB() *DB {
	return &DB{
		Data: make(map[string]string),
		Ttl:  make(map[string]time.Time),
	}
}

func (db *DB) Get(key string) string {
	db.mu.RLock()
	defer db.mu.RUnlock()

	expiry, ok := db.Ttl[key]

	if !ok {
		return "KEY_NOT_FOUND"
	}

	if time.Now().After(expiry) {
		return "Key expired"
	}

	val, ok := db.Data[key]

	return val
}

func (db *DB) Set(key, value string) string {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.Data[key] = value
	db.Ttl[key] = time.Now().Add(config.TimeToLive)

	//log.Print(db.Data)

	return "200 - Value set for Key"
}

func (db *DB) Delete(key string) string {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, ok := db.Data[key]

	if !ok {
		return "KEY_NOT_FOUND"
	}

	delete(db.Data, key)
	delete(db.Ttl, key)

	return "200 - Value deleted for Key"
}

func (db *DB) PeriodicCleaning() {

	for {
		time.Sleep(time.Minute)

		db.mu.Lock()

		now := time.Now()

		for k, v := range db.Ttl {
			if now.After(v) {
				delete(db.Data, k)
				delete(db.Ttl, k)
			}
		}

		db.mu.Unlock()
	}

}

func (db *DB) SaveRDB(path string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	file, err := os.Create(path)

	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(db)
	if err != nil {
		return err
	}

	log.Print("RDB Backup created")

	return nil
}

func (db *DB) LoadRDB(path string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&db)

	return err
}
