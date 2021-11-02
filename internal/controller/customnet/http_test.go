package customnet

import (
	"encoding/json"
	"search_engine/internal/service/engine"
	"search_engine/internal/service/objs"
	"search_engine/internal/util/log"
	"testing"
	"time"
)

func TestGetAddDocArgs(t *testing.T) {
	level := "debug"
	filePath := "/Users/wengguan/search_code/search_file/logs/engine.log"
	maxSize := 128
	maxBackups := 100
	maxAge := 60
	compress := true
	log.InitLogger(level, filePath, maxSize, maxBackups, maxAge, compress)

	var docReq objs.Doc
	docReq = objs.Doc{Ident: "88.199.1/fff.def", Data: objs.Data{Modified: "北京市首都机场", Saled: "成都", CreatedAt: time.Now(), Num: 5}}
	docByte, _ := json.Marshal(&docReq)
	t.Log(string(docByte))
}

func TestGetRetriveArgs(t *testing.T) {
	var rr RetreiveReq
	rr.RetreiveTerms = []objs.RetreiveTerm{{"Modified", "北京", objs.Eq, objs.Union}, {"Num", 12, objs.Gt, objs.Filter}}
	rrByte, _ := json.Marshal(&rr)
	t.Log(string(rrByte))
}

func TestAll(t *testing.T) {
	analyzerStopWordPath := "/Users/wengguan/search_code/search/search_engine/configs/stop_word.txt"
	dbPath := "/Users/wengguan/search_code/search_file/db/engine.db"
	dbHost := "127.0.0.1"
	dbPort := "4379"
	dbAuth := ""
	dbIndex := 0
	dbTimeout := 30
	bloomfilterMiscalRate := 0.00001
	var bloomfilterAddSize uint64 = 100000000
	engine.NewEg(analyzerStopWordPath, dbPath, dbHost, dbPort, dbAuth, dbIndex, dbTimeout, bloomfilterMiscalRate, bloomfilterAddSize)
	ip, port := "", "7788"
	cn := NetFactory("http")
	cn.StartNet(ip, port)
	c := make(chan int)
	_ = <-c
	engine.CloseEg()
}

/*
GET 127.0.0.1:7788/del_doc?docid=0
GET 127.0.0.1:7788/doc_isdel?docid=3
POST 127.0.0.1:7788/add_doc
{"Ident":"88.199.1/fff.def","Modified":"北京市首都机场","Saled":"成都","Num":5,"CreatedAt":"2021-11-02T11:42:08.995206+08:00"}
POST 127.0.0.1:7788/retrieve
{"RetreiveTerms":[{"fieldName":"Modified","term":"北京","TermType":1,"operator":"or"},{"fieldName":"Num","term":12,"TermType":16,"operator":"filter"}]}
*/
