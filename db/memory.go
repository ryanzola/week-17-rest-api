package db

import (
	"fmt"
	"sync"
)

type MemoryRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MemoryDB struct {
	mu sync.Mutex // always put your mutex above the value you want to protect
	db map[string]string
}

func NewMemory() *MemoryDB {
	return &MemoryDB{
		db: make(map[string]string),
	}
}

func (db *MemoryDB) Add(record MemoryRecord) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if record.Key == "" || record.Value == "" {
		return fmt.Errorf("invalid record")
	}

	db.db[record.Key] = record.Value
	return nil
}

func (db *MemoryDB) Get(key string) (MemoryRecord, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if value, ok := db.db[key]; ok {
		record := MemoryRecord{
			Key:   key,
			Value: value,
		}

		return record, true
	}

	return MemoryRecord{}, false
}
