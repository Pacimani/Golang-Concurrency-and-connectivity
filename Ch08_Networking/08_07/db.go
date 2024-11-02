package main

import (
	"sync"
)

// Database
type DB struct {
	mu sync.Map
}

func (db *DB) Set(key string, value []byte) {
	db.mu.Store(key, value)
}

func (db *DB) Get(key string) []byte {

	value, ok := db.mu.Load(key)
	if !ok {
		return nil
	}
	return value.([]byte)
}
