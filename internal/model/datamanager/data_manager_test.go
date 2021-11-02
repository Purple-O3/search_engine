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
	posting := objs.Posting{FieldName: "Modified", Term: "丰台区", Docid: docid}
	ps = append(ps, posting)
	posting = objs.Posting{FieldName: "Saled", Term: "海淀区", Docid: docid}
	ps = append(ps, posting)
	doc := objs.Doc{Ident: "88.199.1/aaa.def", Data: objs.Data{Modified: "北京市丰台区", Saled: "北京市海淀区"}}
	mg.AddDoc(doc, docid, ps)

	docid += 1
	ps = make(objs.Postings, 0)
	posting = objs.Posting{FieldName: "Modified", Term: "黄浦区", Docid: docid}
	ps = append(ps, posting)
	posting = objs.Posting{FieldName: "Saled", Term: "浦东新区", Docid: docid}
	ps = append(ps, posting)
	doc = objs.Doc{Ident: "88.199.1/bbb.def", Data: objs.Data{Modified: "上海市黄浦区", Saled: "上海市浦东新区"}}
	mg.AddDoc(doc, docid, ps)

	docid += 1
	ps = make(objs.Postings, 0)
	posting = objs.Posting{FieldName: "Modified", Term: "河东区", Docid: docid}
	ps = append(ps, posting)
	posting = objs.Posting{FieldName: "Saled", Term: "河西区", Docid: docid}
	ps = append(ps, posting)
	doc = objs.Doc{Ident: "88.199.1/ccc.def", Data: objs.Data{Modified: "天津市河东区", Saled: "天津市河西区"}}
	mg.AddDoc(doc, docid, ps)

	docid += 1
	ps = make(objs.Postings, 0)
	posting = objs.Posting{FieldName: "Modified", Term: "南昌市", Docid: docid}
	ps = append(ps, posting)
	posting = objs.Posting{FieldName: "Saled", Term: "吉安市", Docid: docid}
	ps = append(ps, posting)
	doc = objs.Doc{Ident: "88.199.1/ddd.def", Data: objs.Data{Modified: "江西省南昌市", Saled: "江西省吉安市"}}
	mg.AddDoc(doc, docid, ps)

	trackid := uint64(rand.Intn(999) + 1)
	ctx := context.WithValue(context.Background(), "trackid", trackid)
	ret, _ := mg.Retrieve(ctx, "Modified", "河东区")
	t.Log(ret)
	ret, _ = mg.Retrieve(ctx, "Saled", "河东区")
	t.Log(ret)

	docid += 1
	ps = make(objs.Postings, 0)
	posting = objs.Posting{FieldName: "Modified", Term: "南昌市", Docid: docid}
	ps = append(ps, posting)
	posting = objs.Posting{FieldName: "Saled", Term: "井冈山市", Docid: docid}
	ps = append(ps, posting)
	doc = objs.Doc{Ident: "88.199.1/eee.def", Data: objs.Data{Modified: "江西省南昌市", Saled: "江西省井冈山市"}}
	mg.AddDoc(doc, docid, ps)

	ret, _ = mg.Retrieve(ctx, "Modified", "南昌市")
	t.Log(ret)
}
