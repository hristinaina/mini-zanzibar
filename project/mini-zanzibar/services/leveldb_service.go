package services

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type ILevelDBService interface {
	GetAll() (map[string]string, error)
	GetByKey(key string) (string, error)
	Add(key, value string) error
	Delete(key string) error
}

type LevelDBService struct {
	db *leveldb.DB
}

func NewLevelDBService(db *leveldb.DB) ILevelDBService {
	return &LevelDBService{db}
}

func (ls *LevelDBService) GetAll() (map[string]string, error) {
	iterator := ls.db.NewIterator(nil, nil)
	defer iterator.Release()

	data := make(map[string]string)
	for iterator.Next() {
		key := iterator.Key()
		value := iterator.Value()
		data[string(key)] = string(value)
	}
	if err := iterator.Error(); err != nil {
		return nil, err
	}

	return data, nil
}

func (ls *LevelDBService) GetByKey(key string) (string, error) {
	value, err := ls.db.Get([]byte(key), nil)
	return string(value), err
}

func (ls *LevelDBService) Add(key, value string) error {
	return ls.db.Put([]byte(key), []byte(value), nil)
}

func (ls *LevelDBService) Delete(key string) error {
	return ls.db.Delete([]byte(key), nil)
}
