package store

import (
	"search_engine/internal/util/log"
	"testing"
)

func TestAll(t *testing.T) {
	level := "debug"
	filePath := "/Users/wengguan/search_code/search_file/logs/engine.log"
	maxSize := 128
	maxBackups := 100
	maxAge := 60
	compress := true
	log.InitLogger(level, filePath, maxSize, maxBackups, maxAge, compress)

	path := "/Users/wengguan/search_code/search_file/db/engine.db"
	host := "192.168.3.4"
	port := "9221"
	auth := ""
	index := 0
	timeout := 30
	s, err := StoreFactory("pika", path, host, port, auth, index, timeout)
	if err != nil {
		t.Log("new s error", err)
		return
	}
	s.Set([]byte("hello"), []byte("hanmeimei"))
	v, err := s.Get([]byte("hello"))
	if err != nil {
		t.Log(err)
	} else {
		t.Log(string(v))
	}
	s.Close()
}

func TestGet(t *testing.T) {
	path := "/Users/wengguan/search_code/search_file/db/engine.db"
	host := "192.168.3.4"
	port := "9221"
	auth := ""
	index := 0
	timeout := 30
	s, err := StoreFactory("pika", path, host, port, auth, index, timeout)
	if err != nil {
		t.Log("new s error", err)
		return
	}
	v, err := s.Get([]byte("hello"))
	if err != nil {
		t.Log(err)
	} else {
		t.Log(string(v))
	}
	s.Delete([]byte("hello"))
	v, err = s.Get([]byte("hello"))
	if err != nil {
		t.Log(err)
	} else {
		t.Log(string(v))
	}
	s.Close()
}
