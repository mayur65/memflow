package tests

import (
	"github.com/mayur65/memflow/internal/storage"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestSetGet(t *testing.T) {
	db := storage.InitDB()

	response := db.Set("Key", "Value")
	assert.Equal(t, "200 - Value set for Key", response, "SET response mismatch")

	value := db.Get("Key")
	assert.Equal(t, "Value", value, "GET response mismatch")
}

func TestKeyNotFound(t *testing.T) {
	db := storage.InitDB()

	value := db.Get("Key")
	assert.Equal(t, "KEY_NOT_FOUND", value, "GET response mismatch")
}

func TestKeyDelete(t *testing.T) {
	db := storage.InitDB()

	response := db.Set("Key", "Value")
	assert.Equal(t, "200 - Value set for Key", response, "SET response mismatch")

	response = db.Delete("Key")
	assert.Equal(t, "200 - Value deleted for Key", response, "DELETE response mismatch")
}

func TestKeyExpired(t *testing.T) {
	db := storage.InitDB()

	db.Set("Key", "Value")

	time.Sleep(30 * time.Second)

	value := db.Get("Key")
	assert.Equal(t, "Key expired", value, "GET response mismatch")
}

func TestKeyCleanup(t *testing.T) {
	db := storage.InitDB()

	go db.PeriodicCleaning()

	db.Set("Key", "Value")

	time.Sleep(65 * time.Second)

	value := db.Get("Key")
	assert.Equal(t, "KEY_NOT_FOUND", value, "GET response mismatch")
}

func TestRDBSaved(t *testing.T) {
	// Create db, set kv pair and save RDB snapshot.
	db := storage.InitDB()
	db.Set("Key", "Value")
	err := db.SaveRDB("test.rdb")
	assert.NoError(t, err, "Error saving RDB.")

	// Create newDb, load the saved snapshot, and assert on kv pair.
	newDb := storage.InitDB()
	err = newDb.LoadRDB("test.rdb")
	assert.NoError(t, err, "Error loading RDB.")
	value := newDb.Get("Key")
	assert.Equal(t, "Value", value, "GET response mismatch")

	_ = os.Remove("test.rdb")
}
