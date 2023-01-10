package store

import (
	"search_engine/internal/objs"
	"testing"
)

func TestAll(t *testing.T) {
	config := objs.DBConfig{Type: "pika", Path: "../../../data/db/engine.db", Host: "192.168.3.4", Password: "", Index: 0, Timeout: 30}
	s, err := StoreFactory(config)
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
	config := objs.DBConfig{Type: "pika", Path: "../../../data/db/engine.db", Host: "192.168.3.4", Password: "", Index: 0, Timeout: 30}
	s, err := StoreFactory(config)
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
