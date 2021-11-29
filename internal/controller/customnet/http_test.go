package customnet

import (
	"encoding/json"
	"search_engine/internal/service/engine"
	"search_engine/internal/service/objs"
	"search_engine/internal/util/bloomfilter"
	"search_engine/internal/util/log"
	"testing"
	"time"
)

func TestGetAddDocArgs(t *testing.T) {
	level := "debug"
	filePath := "../../../logs/engine.log"
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
	var rr objs.RetreiveReq
	rr.RetreiveTerms = []objs.RetreiveTerm{{FieldName: "Modified", Term: "北京", TermCompareType: objs.Eq, Operator: objs.Union}, {FieldName: "Ident", Term: "88.199.1/fff.def", TermCompareType: objs.Eq, Operator: objs.Union}, {FieldName: "Num", Term: 12, TermCompareType: objs.Gt, Operator: objs.Filter}, {FieldName: "CreatedAt", Term: "2021-11-03T15:14:05.126975+08:00", TermCompareType: objs.Lt, Operator: objs.Filter}}
	rrByte, _ := json.Marshal(&rr)
	t.Log(string(rrByte))
}

func TestAll(t *testing.T) {
	analyzerStopWordPath := "../../../data/stop_word.txt"
	dbPath := "../../../data/db/engine.db"
	dbHost := "127.0.0.1"
	dbPort := "4379"
	dbAuth := ""
	dbIndex := 0
	dbTimeout := 30
	bloomfilterMiscalRate := 0.00001
	var bloomfilterAddSize uint64 = 100000000
	bloomfilterStorePath := "../../../data/bloomfilter"
	engine.NewEg(analyzerStopWordPath, dbPath, dbHost, dbPort, dbAuth, dbIndex, dbTimeout, bloomfilterMiscalRate, bloomfilterAddSize, bloomfilterStorePath)
	ip, port := "", "7788"
	cn := NetFactory("http")
	cn.StartNet(ip, port)
	/*c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c*/
	engine.CloseEg()
	log.CloseLogger()
	cn.Shutdown()
	bloomfilter.DeleteBloomFile(bloomfilterStorePath)
}

/*
GET 127.0.0.1:7788/del_doc?docid=0
GET 127.0.0.1:7788/doc_isdel?docid=3
POST 127.0.0.1:7788/add_doc
{"Ident":"88.199.1/bbb.def","Modified":"北京市首都机场","Saled":"成都","Num":13,"CreatedAt":"2021-11-02T16:42:21.199502+08:00"}
POST 127.0.0.1:7788/retrieve
{"RetreiveTerms":[{"FieldName":"Modified","Term":"北京","TermCompareType":1,"Operator":"must"},{"FieldName":"Num","Term":12,"TermCompareType":16,"Operator":"filter"}],"Offset":0,"Limit":10}
*/
