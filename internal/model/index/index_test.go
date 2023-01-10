package index

import (
	"search_engine/internal/model/store"
	"search_engine/internal/objs"
	"strconv"
	"testing"
)

func TestAll(t *testing.T) {
	config := objs.DBConfig{Type: "pika", Path: "../../../data/db/engine.db", Host: "192.168.3.4", Index: 0, Timeout: 30}
	s, err := store.StoreFactory(config)
	if err != nil {
		t.Log("new s error", err)
		return
	}

	invertedIndex := NewInvertedIndex(s)
	var docid uint64 = 0
	posting := objs.Posting{FieldName: "context", Term: "唐时", Docid: docid}
	invertedIndex.Set(posting.FieldName+"_"+posting.Term, posting)
	docid = 1
	posting = objs.Posting{FieldName: "context", Term: "明月", Docid: docid}
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
