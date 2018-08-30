package leveldb

import (
	"github.com/republicprotocol/renex-sdk-go/adapter/store"
	"github.com/syndtr/goleveldb/leveldb"
)

type ldbStore struct {
	db *leveldb.DB
}

type Store interface {
	Read(key []byte) ([]byte, error)
	Write(key []byte, value []byte) error
	Delete(key []byte) error
	Close() error
}

func NewLDBStore(path string) (Store, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &ldbStore{
		db: db,
	}, nil
}

func (ldb *ldbStore) Read(key []byte) ([]byte, error) {
	value, err := ldb.db.Get(key, nil)
	if err == leveldb.ErrNotFound {
		return nil, store.ErrOrdersNotFound
	}
	return value, err
}

func (ldb *ldbStore) Write(key []byte, value []byte) error {
	return ldb.db.Put(key, value, nil)
}

func (ldb *ldbStore) Delete(key []byte) error {
	return ldb.db.Delete(key, nil)
}

func (ldb *ldbStore) Close() error {
	return ldb.db.Close()
}
