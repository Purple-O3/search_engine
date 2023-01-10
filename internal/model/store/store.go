package store

import (
	"errors"
	"search_engine/internal/objs"
	"search_engine/internal/util/log"
	"search_engine/internal/util/rediswrapper"
)

type Store interface {
	Set(k []byte, v []byte) error
	Get(k []byte) ([]byte, error)
	Delete(k []byte) error
	Close() error
}

func StoreFactory(config objs.DBConfig) (Store, error) {
	switch config.Type {
	case "pika":
		return rediswrapper.NewRedis(config.Host, config.Port, config.Password, config.Index, config.Timeout)
	//case "rocksdb":
	//	return rocksdbwrapper.NewRocksdb(path)
	default:
		log.Errorf("store_type argv error")
		return nil, errors.New("store_type argv error")
	}
}
