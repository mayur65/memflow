package storage

import (
	"fmt"
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

	//expiry, ok := db.ttl[key]
	//
	//if !ok {
	//	return "Key does not exist"
	//}
	//
	//if time.Now().After(expiry) {
	//	return "Key expired"
	//}

	val, ok := db.data[key]
	fmt.Println(key)
	fmt.Println(ok)
	fmt.Println(key)
	if !ok {
		return "KEY_NOT_FOUND"
	}

	return val
}

func (db *DB) Set(key, value string) string {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[key] = value

	//if ttl > 0 {
	//	db.ttl[key] = time.Now().Add(ttl)
	//} else {
	//	delete(db.ttl, key)
	//}

	fmt.Println(db.data)

	return "200 - Value set for Key"
}

//Implement periodic expired kv clean-up
