package datamanager

import (
	"context"
	"math/rand"
	"search_engine/internal/service/objs"
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

	dbPath := "/Users/wengguan/search_code/search_file/db/engine.db"
	dbHost := "127.0.0.1"
	dbPort := "4379"
	dbAuth := ""
	dbIndex := 0
	dbTimeout := 30
	mg := NewManager(dbPath, dbHost, dbPort, dbAuth, dbIndex, dbTimeout)

	var docid uint64 = 0
	ps := make(objs.Postings, 0)
	posting := objs.Posting{Term: "浪漫", Docid: docid, TermFreq: 5, Offset: []int{1640, 1034, 1223, 1491, 1678}}
	ps = append(ps, posting)
	posting = objs.Posting{Term: "巴黎", Docid: docid, TermFreq: 2, Offset: []int{5048, 6813}}
	ps = append(ps, posting)
	posting = objs.Posting{Term: "土耳其", Docid: docid, TermFreq: 1, Offset: []int{4247}}
	ps = append(ps, posting)
	doc := objs.Doc{Body: "浪漫巴黎土耳其", Title: "五零班", Price: 5.00}
	mg.AddDoc(doc, docid, ps)

	docid += 1
	ps = make(objs.Postings, 0)
	posting = objs.Posting{Term: "明朝", Docid: docid, TermFreq: 2, Offset: []int{1640, 1447}}
	ps = append(ps, posting)
	posting = objs.Posting{Term: "那些", Docid: docid, TermFreq: 3, Offset: []int{5048, 6813, 1644}}
	ps = append(ps, posting)
	posting = objs.Posting{Term: "事儿", Docid: docid, TermFreq: 4, Offset: []int{3692, 4312, 5115, 5116}}
	ps = append(ps, posting)
	doc = objs.Doc{Body: "明朝那些事儿", Title: "五一班", Price: 5.10}
	mg.AddDoc(doc, docid, ps)

	docid += 1
	ps = make(objs.Postings, 0)
	posting = objs.Posting{Term: "银河", Docid: docid, TermFreq: 3, Offset: []int{1640, 1914, 1938}}
	ps = append(ps, posting)
	posting = objs.Posting{Term: "英雄", Docid: docid, TermFreq: 6, Offset: []int{5048, 6813, 5019, 5074, 6339, 6946}}
	ps = append(ps, posting)
	posting = objs.Posting{Term: "传说", Docid: docid, TermFreq: 2, Offset: []int{4247, 6868}}
	ps = append(ps, posting)
	doc = objs.Doc{Body: "银河英雄传说", Title: "五二班", Price: 5.20}
	mg.AddDoc(doc, docid, ps)

	docid += 1
	ps = make(objs.Postings, 0)
	posting = objs.Posting{Term: "万里", Docid: docid, TermFreq: 3, Offset: []int{5048, 6813, 1678}}
	ps = append(ps, posting)
	posting = objs.Posting{Term: "长城", Docid: docid, TermFreq: 4, Offset: []int{4247, 746, 778, 3110}}
	ps = append(ps, posting)
	doc = objs.Doc{Body: "中国万里长城", Title: "五三班", Price: 5.30}
	mg.AddDoc(doc, docid, ps)

	trackid := rand.Intn(999) + 1
	ctx := context.WithValue(context.Background(), "trackid", trackid)
	ret, _ := mg.Retrieve(ctx, "英雄")
	t.Log(ret)
	ret, _ = mg.Retrieve(ctx, "万里")
	t.Log(ret)

	docid += 1
	ps = make(objs.Postings, 0)
	posting = objs.Posting{Term: "埃及", Docid: docid, TermFreq: 1, Offset: []int{5123}}
	ps = append(ps, posting)
	posting = objs.Posting{Term: "金字塔", Docid: docid, TermFreq: 2, Offset: []int{4320, 756}}
	ps = append(ps, posting)
	doc = objs.Doc{Body: "埃及金字塔", Title: "五四班", Price: 5.40}
	mg.AddDoc(doc, docid, ps)

	ret, _ = mg.Retrieve(ctx, "埃及")
	t.Log(ret)
}
