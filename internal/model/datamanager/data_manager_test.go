package datamanager

import (
	"math/rand"
	"search_engine/internal/objs"
	"testing"
)

func TestAll(t *testing.T) {
	config := objs.DBConfig{Type: "pika", Path: "../../../data/db/engine.db", Host: "${DBHOST||localhost}", Port: 9221, Password: "", Index: 0, Timeout: 30}
	mg := NewManager(config)

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
	ret, _ := mg.Retrieve("Modified", "河东区", trackid)
	t.Log(ret)
	ret, _ = mg.Retrieve("Saled", "河东区", trackid)
	t.Log(ret)

	docid += 1
	ps = make(objs.Postings, 0)
	posting = objs.Posting{FieldName: "Modified", Term: "南昌市", Docid: docid}
	ps = append(ps, posting)
	posting = objs.Posting{FieldName: "Saled", Term: "井冈山市", Docid: docid}
	ps = append(ps, posting)
	doc = objs.Doc{Ident: "88.199.1/eee.def", Data: objs.Data{Modified: "江西省南昌市", Saled: "江西省井冈山市"}}
	mg.AddDoc(doc, docid, ps)

	ret, _ = mg.Retrieve("Modified", "南昌市", trackid)
	t.Log(ret)
}
