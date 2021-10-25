package engine

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

	analyzerStopWordPath := "/Users/wengguan/search_code/search/search_engine/configs/stop_word.txt"
	dbPath := "/Users/wengguan/search_code/search_file/db/engine.db"
	dbHost := "127.0.0.1"
	dbPort := "4379"
	dbAuth := ""
	dbIndex := 0
	dbTimeout := 30
	bloomfilterMiscalRate := 0.00001
	var bloomfilterAddSize uint64 = 100000000
	engine := newEngine(analyzerStopWordPath, dbPath, dbHost, dbPort, dbAuth, dbIndex, dbTimeout, bloomfilterMiscalRate, bloomfilterAddSize)
	defer engine.close()

	var docid uint64 = 0
	doc := objs.Doc{Body: "浪漫巴黎土耳其", Title: "五零班", Price: 5.00}
	trackid := rand.Intn(999) + 1
	ctx := context.WithValue(context.Background(), "trackid", trackid)
	engine.addDoc(ctx, doc, docid)

	docid = 1
	doc = objs.Doc{Body: "明朝那些事儿", Title: "五一班", Price: 5.10}
	engine.addDoc(ctx, doc, docid)

	docid = 2
	doc = objs.Doc{Body: "银河英雄传说", Title: "五三班", Price: 5.20}
	engine.addDoc(ctx, doc, docid)

	docid = 3
	doc = objs.Doc{Body: "中国万里长城", Title: "五三班", Price: 5.30}
	engine.addDoc(ctx, doc, docid)

	docid = 4
	doc = objs.Doc{Body: "埃及金字塔", Title: "五四班", Price: 5.40}
	engine.addDoc(ctx, doc, docid)

	retreiveTerms := []objs.RetreiveTerm{{"英雄", true, false}, {"埃及", true, false}, {"长城", true, false}}
	titleMust := "五三班"
	priceStart := 5.10
	priceEnd := 5.50
	ret := engine.retrieveDoc(ctx, retreiveTerms, titleMust, priceStart, priceEnd)
	t.Logf("%v", ret)

	engine.delDoc(0)
	ret = engine.retrieveDoc(ctx, retreiveTerms, titleMust, priceStart, priceEnd)
	t.Logf("%v", ret)

	engine.delDoc(3)
	ret = engine.retrieveDoc(ctx, retreiveTerms, titleMust, priceStart, priceEnd)
	t.Logf("%v", ret)
}

func TestCalInter(t *testing.T) {
	analyzerStopWordPath := "/Users/wengguan/search_code/search/search_engine/configs/stop_word.txt"
	dbPath := "/Users/wengguan/search_code/search_file/db/engine.db"
	dbHost := "127.0.0.1"
	dbPort := "4379"
	dbAuth := ""
	dbIndex := 0
	dbTimeout := 30
	bloomfilterMiscalRate := 0.00001
	var bloomfilterAddSize uint64 = 100000000
	engine := newEngine(analyzerStopWordPath, dbPath, dbHost, dbPort, dbAuth, dbIndex, dbTimeout, bloomfilterMiscalRate, bloomfilterAddSize)
	defer engine.close()

	replUniqInters := make([]objs.RecallPostingList, 0)
	replUniqInter := make(objs.RecallPostingList, 0)
	reposting := objs.RecallPosting{objs.Posting{Term: "埃及", Docid: 3, TermFreq: 1, Offset: []int{4}}, objs.Doc{Body: "", Title: "五三班", Price: 5.300000}}
	replUniqInter = append(replUniqInter, reposting)
	reposting = objs.RecallPosting{objs.Posting{Term: "埃及", Docid: 4, TermFreq: 1, Offset: []int{0}}, objs.Doc{Body: "", Title: "五四班", Price: 5.400000}}
	replUniqInter = append(replUniqInter, reposting)
	replUniqInters = append(replUniqInters, replUniqInter)

	replUniqInter = make(objs.RecallPostingList, 0)
	reposting = objs.RecallPosting{objs.Posting{Term: "银河", Docid: 1, TermFreq: 1, Offset: []int{0}}, objs.Doc{Body: "", Title: "五二班", Price: 5.200000}}
	replUniqInter = append(replUniqInter, reposting)
	reposting = objs.RecallPosting{objs.Posting{Term: "银河", Docid: 2, TermFreq: 1, Offset: []int{0}}, objs.Doc{Body: "", Title: "五三班", Price: 5.300000}}
	replUniqInter = append(replUniqInter, reposting)
	reposting = objs.RecallPosting{objs.Posting{Term: "银河", Docid: 4, TermFreq: 1, Offset: []int{0}}, objs.Doc{Body: "", Title: "五三班", Price: 5.300000}}
	replUniqInter = append(replUniqInter, reposting)
	replUniqInters = append(replUniqInters, replUniqInter)

	replUniqUnion := make(objs.RecallPostingList, 0)
	reposting = objs.RecallPosting{objs.Posting{Term: "中国", Docid: 1, TermFreq: 1, Offset: []int{0}}, objs.Doc{Body: "", Title: "五三班", Price: 5.300000}}
	replUniqUnion = append(replUniqUnion, reposting)
	reposting = objs.RecallPosting{objs.Posting{Term: "明朝", Docid: 2, TermFreq: 1, Offset: []int{0}}, objs.Doc{Body: "", Title: "五四班", Price: 5.400000}}
	replUniqUnion = append(replUniqUnion, reposting)
	reposting = objs.RecallPosting{objs.Posting{Term: "银河", Docid: 4, TermFreq: 1, Offset: []int{4}}, objs.Doc{Body: "", Title: "五四班", Price: 5.400000}}
	replUniqUnion = append(replUniqUnion, reposting)
	reposting = objs.RecallPosting{objs.Posting{Term: "埃及", Docid: 5, TermFreq: 1, Offset: []int{4}}, objs.Doc{Body: "", Title: "五四班", Price: 5.400000}}
	replUniqUnion = append(replUniqUnion, reposting)

	ret := engine.calInter(replUniqUnion, replUniqInters)
	t.Logf("%v", ret)
}

func TestWrap(t *testing.T) {
	analyzerStopWordPath := "/Users/wengguan/search_code/search/search_engine/configs/stop_word.txt"
	dbPath := "/Users/wengguan/search_code/search_file/db/engine.db"
	dbHost := "127.0.0.1"
	dbPort := "4379"
	dbAuth := ""
	dbIndex := 0
	dbTimeout := 30
	bloomfilterMiscalRate := 0.00001
	var bloomfilterAddSize uint64 = 100000000
	NewEg(analyzerStopWordPath, dbPath, dbHost, dbPort, dbAuth, dbIndex, dbTimeout, bloomfilterMiscalRate, bloomfilterAddSize)
	defer CloseEg()

	trackid := rand.Intn(999) + 1
	ctx := context.WithValue(context.Background(), "trackid", trackid)
	doc := objs.Doc{Body: "浪漫巴黎土耳其", Title: "五零班", Price: 5.00}
	AddDoc(ctx, doc)

	doc = objs.Doc{Body: "明朝那些事儿", Title: "五一班", Price: 5.10}
	AddDoc(ctx, doc)

	doc = objs.Doc{Body: "银河英雄传说", Title: "五三班", Price: 5.20}
	AddDoc(ctx, doc)

	doc = objs.Doc{Body: "中国万里长城", Title: "五三班", Price: 5.30}
	AddDoc(ctx, doc)

	doc = objs.Doc{Body: "埃及金字塔", Title: "五四班", Price: 5.40}
	AddDoc(ctx, doc)

	retreiveTerms := []objs.RetreiveTerm{{"英雄", true, false}, {"埃及", true, false}, {"长城", true, false}}
	titleMust := "五三班"
	priceStart := 5.10
	priceEnd := 5.50
	ret := RetrieveDoc(ctx, retreiveTerms, titleMust, priceStart, priceEnd)
	t.Logf("%v", ret)
}

func TestAll2(t *testing.T) {
	analyzerStopWordPath := "/Users/wengguan/search_code/search/search_engine/configs/stop_word.txt"
	dbPath := "/Users/wengguan/search_code/search_file/db/engine.db"
	dbHost := "127.0.0.1"
	dbPort := "4379"
	dbAuth := ""
	dbIndex := 0
	dbTimeout := 30
	bloomfilterMiscalRate := 0.00001
	var bloomfilterAddSize uint64 = 100000000
	engine := newEngine(analyzerStopWordPath, dbPath, dbHost, dbPort, dbAuth, dbIndex, dbTimeout, bloomfilterMiscalRate, bloomfilterAddSize)
	defer engine.close()

	trackid := rand.Intn(999) + 1
	ctx := context.WithValue(context.Background(), "trackid", trackid)
	var docid uint64 = 0
	doc := objs.Doc{Body: "浪漫巴黎土耳其", Title: "五零班", Price: 5.00}
	engine.addDoc(ctx, doc, docid)

	docid = 1
	doc = objs.Doc{Body: "明朝那些事儿", Title: "五一班", Price: 5.10}
	engine.addDoc(ctx, doc, docid)

	docid = 2
	doc = objs.Doc{Body: "银河英雄传说", Title: "五三班", Price: 5.20}
	engine.addDoc(ctx, doc, docid)

	docid = 3
	doc = objs.Doc{Body: "中国万里长城", Title: "五三班", Price: 5.30}
	engine.addDoc(ctx, doc, docid)

	docid = 4
	doc = objs.Doc{Body: "埃及金字塔", Title: "五四班", Price: 5.40}
	engine.addDoc(ctx, doc, docid)

	retreiveTerms := []objs.RetreiveTerm{{"英雄", true, false}, {"埃及", true, false}, {"长城", true, false}}
	titleMust := "五三班"
	priceStart := 5.10
	priceEnd := 5.50
	ret := engine.retrieveDoc(ctx, retreiveTerms, titleMust, priceStart, priceEnd)
	t.Logf("%v", ret)

	engine.delDoc(0)
	ret = engine.retrieveDoc(ctx, retreiveTerms, titleMust, priceStart, priceEnd)
	t.Logf("%v", ret)

	engine.delDoc(3)
	ret = engine.retrieveDoc(ctx, retreiveTerms, titleMust, priceStart, priceEnd)
	t.Logf("%v", ret)
}
