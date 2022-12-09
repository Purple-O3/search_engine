package store

import (
	"testing"
)

func TestAll(t *testing.T) {
	path := "../../../data/db/engine.db"
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
	path := "../../../data/db/engine.db"
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
