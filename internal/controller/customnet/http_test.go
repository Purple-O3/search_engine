package customnet

import (
	"encoding/json"
	"search_engine/internal/service/engine"
	"search_engine/internal/service/objs"
	"search_engine/internal/util/log"
	"testing"
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
	docReq = objs.Doc{Body: "银河英雄传说", Title: "五三班", Price: 5.20}
	docByte, _ := json.Marshal(&docReq)
	t.Log(string(docByte))
}

func TestGetRetriveArgs(t *testing.T) {
	var rr RetreiveReq
	rr.RetreiveTerms = []objs.RetreiveTerm{{"英雄", "union"}, {"埃及", "union"}, {"长城", "inter"}}
	rr.TitleMust = "五三班"
	rr.PriceStart = 5.10
	rr.PriceEnd = 5.50
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
{"body": "浪漫巴黎土耳其", "title": "五零班", "price": 5.00}
{"body": "明朝那些事儿", "title": "五一班", "price": 5.10}
{"body": "银河英雄传说", "title": "五二班", "price": 5.20}
{"body": "中国万里长城", "title": "五三班", "price": 5.30}
{"body": "埃及金字塔", "title": "五四班", "price": 5.40}
POST 127.0.0.1:7788/retrieve
{"retreive_terms":[{"term":"英雄","op_type":"union"},{"term":"埃及","op_type":"union"},{"term":"长城","op_type":"inter"}],"title_must":"五三班","price_start":5.1,"price_end":5.5}
*/
