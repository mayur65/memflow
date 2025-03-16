package tests

import (
	"github.com/mayur65/memflow/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
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
