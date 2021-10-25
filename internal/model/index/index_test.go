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
	s, err := store.StoreFactory("pika", path, host, port, auth, index, timeout)
	if err != nil {
		t.Log("new s error", err)
		return
	}

	invertedIndex := NewInvertedIndex(s)
	var docid uint64 = 0
	posting := objs.Posting{"唐时", docid, 2, []int{4, 6}}
	invertedIndex.Set(posting.Term, posting)
	docid = 1
	posting = objs.Posting{"明月", docid, 1, []int{7}}
	invertedIndex.Set(posting.Term, posting)
	pl, _ := invertedIndex.Get(posting.Term)
	t.Log(pl)

	positiveIndex := NewPositiveIndex(s)
	docid = 2
	docidString := strconv.FormatUint(docid, 10)
	docKey := docidString + "_text"
	text := "汉时"
	positiveIndex.Set(docKey, text)
	docid = 3
	text = "关"
	docidString = strconv.FormatUint(docid, 10)
	docKey = docidString + "_text"
	positiveIndex.Set(docKey, text)
	value, _ := positiveIndex.Get(docKey)
	t.Log(value)
}
