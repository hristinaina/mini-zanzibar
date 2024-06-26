package config

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

func InitLevelDB() *leveldb.DB {
	db, err := leveldb.OpenFile("/databases/leveldb", nil)
	if err != nil {
		log.Fatalf("Failed to open LevelDB: %v", err)
	}
	return db
}

func CloseLevelDB(db *leveldb.DB) {
	if err := db.Close(); err != nil {
		log.Fatalf("Failed to close LevelDB: %v", err)
	}
}
