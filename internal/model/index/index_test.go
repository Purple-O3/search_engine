package index

import (
	"search_engine/internal/model/store"
	"search_engine/internal/service/objs"
	"search_engine/internal/util/log"
	"strconv"
	"testing"
)

func TestAll(t *testing.T) {
	level := "debug"
	filePath := "../../../logs/engine.log"
	maxSize := 128
	maxBackups := 100
	maxAge := 60
	compress := true
	log.InitLogger(level, filePath, maxSize, maxBackups, maxAge, compress)

	path := "../../../data/db/engine.db"
	host := "192.168.3.4"
	port := "9221"
	auth := ""
	index := 0
	timeout := 30
	s, err := store.StoreFactory("pika", path, host, port, auth, index, timeout)
	if err != nil {
		t.Log("new s error", err)
		return
	}

	invertedIndex := NewInvertedIndex(s)
	var docid uint64 = 0
	posting := objs.Posting{"context", "唐时", docid}
	invertedIndex.Set(posting.FieldName+"_"+posting.Term, posting)
	docid = 1
	posting = objs.Posting{"context", "明月", docid}
	invertedIndex.Set(posting.FieldName+"_"+posting.Term, posting)
	pl, _ := invertedIndex.Get(posting.FieldName + "_" + posting.Term)
	t.Log(pl)

	positiveIndex := NewPositiveIndex(s)
	docid = 2
	docidString := strconv.FormatUint(docid, 10)
	docKey := "doc" + docidString
	text := "汉时"
	positiveIndex.Set(docKey, text)
	docid = 3
	text = "关"
	docidString = strconv.FormatUint(docid, 10)
	docKey = "doc" + docidString
	positiveIndex.Set(docKey, text)
	value, _ := positiveIndex.Get(docKey)
	t.Log(value)
}
