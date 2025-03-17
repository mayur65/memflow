package tests

import (
	"github.com/mayur65/memflow/internal/storage"
	"github.com/stretchr/testify/assert"
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
