package store

import (
	"errors"
	"search_engine/internal/util/log"
	"search_engine/internal/util/wrapredis"
)

type Store interface {
	Set(k []byte, v []byte) error
	Get(k []byte) ([]byte, error)
	Delete(k []byte) error
	Close() error
}

func StoreFactory(storeType string, path string, host string, port string, password string, index int, timeout int) (Store, error) {
	switch storeType {
	//case "rocksdb":
	//	return wraprocksdb.NewRocksdb(path)
	case "pika":
		return wrapredis.NewRedis(host, port, password, index, timeout)
		//case "fileSystem":
		//		return NewFileSystem(dbPath)
	default:
		log.Errorf("store_type argv error")
		return nil, errors.New("store_type argv error")
	}
}
