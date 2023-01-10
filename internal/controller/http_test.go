package controller

import (
	"encoding/json"
	"search_engine/internal/objs"
	enginepack "search_engine/internal/service/engine"
	"search_engine/internal/util/bloomfilter"
	"search_engine/internal/util/log"
	"testing"
	"time"
)

func TestGetAddDocArgs(t *testing.T) {
	var docReq objs.Doc
	docReq = objs.Doc{Ident: "88.199.1/fff.def", Data: objs.Data{Modified: "北京市首都机场", Saled: "成都", CreatedAt: time.Now(), Num: 5}}
	docByte, _ := json.Marshal(&docReq)
	t.Log(string(docByte))
}

func TestGetRetriveArgs(t *testing.T) {
	var rr objs.RetreiveReq
	rr.RetreiveTerms = []objs.RetreiveTerm{{FieldName: "Modified", Term: "北京", TermCompareType: objs.Eq, Operator: objs.Union}, {FieldName: "Ident", Term: "88.199.1/fff.def", TermCompareType: objs.Eq, Operator: objs.Union}, {FieldName: "Num", Term: 12, TermCompareType: objs.Gt, Operator: objs.Filter}, {FieldName: "CreatedAt", Term: "2021-11-03T15:14:05.126975+08:00", TermCompareType: objs.Lt, Operator: objs.Filter}}
	rrByte, _ := json.Marshal(&rr)
	t.Log(string(rrByte))
}

func TestAll(t *testing.T) {
	config := objs.Config{
		objs.ServerConfig{IP: "127.0.0.1", Port: 7788, Debug: false, ReadTimeout: 1000, WriteTimeout: 1000, IdleTimeout: 1000},
		objs.LogConfig{},
		objs.AnalyzerConfig{StopWordPath: "../../../data/stop_word.txt"},
		objs.BloomfilterConfig{MiscalRate: 0.00001, AddSize: 100000000, StorePath: "../../../data/bloomfilter"},
		objs.DBConfig{Type: "pika", Path: "../../../data/db/engine.db", Host: "${DBHOST||localhost}", Port: 9221, Password: "", Index: 0, Timeout: 30},
	}

	enginepack.NewWrapEngine(config.Analyzer, config.DB, config.Bloomfilter)
	if err := StartNet(config.Server, closeFunc); err != nil {
		panic(err)
	}
	bloomfilter.DeleteBloomFile(config.Bloomfilter.StorePath)
}

func closeFunc() {
	enginepack.CloseWrapEngine()
	log.CloseLogger()
}

/*
GET 127.0.0.1:7788/api/v1/del_doc?docid=0
GET 127.0.0.1:7788/api/v1/doc_isdel?docid=3
POST 127.0.0.1:7788/api/v1/add_doc
{"Ident":"88.199.1/bbb.def","Modified":"北京市首都机场","Saled":"成都","Num":13,"CreatedAt":"2021-11-02T16:42:21.199502+08:00"}
POST 127.0.0.1:7788/api/v1/retrieve_doc
{"RetreiveTerms":[{"FieldName":"Modified","Term":"北京","TermCompareType":1,"Operator":"must"},{"FieldName":"Num","Term":12,"TermCompareType":16,"Operator":"filter"}],"Offset":0,"Limit":10}
*/
