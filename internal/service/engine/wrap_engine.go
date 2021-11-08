package engine

import (
	"context"
	"search_engine/internal/service/objs"
	"search_engine/internal/util/idgenerator"
	"search_engine/internal/util/log"
	"search_engine/internal/util/tools"
	"time"
)

var eg *engine

func NewEg(analyzerStopWordPath string, dbPath string, dbHost string, dbPort string, dbPassword string, dbIndex int, dbTimeout int, bloomfilterMiscalRate float64, bloomfilterAddSize uint64, bloomfilterStorePath string) {
	eg = newEngine(analyzerStopWordPath, dbPath, dbHost, dbPort, dbPassword, dbIndex, dbTimeout, bloomfilterMiscalRate, bloomfilterAddSize, bloomfilterStorePath)
	idgenerator.NewIdGenerator()
}

func CloseEg() {
	eg.close()
}

func RetrieveDoc(ctx context.Context, retreiveTerms []objs.RetreiveTerm) objs.RecallPostingList {
	defer func(cost func() time.Duration) {
		log.Warnf("trackid:%v, cost: %.3f ms", ctx.Value("trackid"), float64(cost().Microseconds())/1000.0)
	}(tools.TimeCost())

	repl := eg.retrieveDoc(ctx, retreiveTerms)
	log.Infof("trackid:%v, RetreiveTerm:%v, repl:%v", ctx.Value("trackid"), retreiveTerms, repl)
	return repl
}

func AddDoc(ctx context.Context, doc objs.Doc) uint64 {
	defer func(cost func() time.Duration) {
		log.Warnf("trackid:%v, cost: %.3f ms", ctx.Value("trackid"), float64(cost().Microseconds())/1000.0)
	}(tools.TimeCost())

	docid := uint64(idgenerator.Generate())
	eg.addDoc(ctx, doc, docid)
	log.Infof("trackid:%v, docid:%d, doc:%v", ctx.Value("trackid"), docid, doc)
	return docid
}

func DelDoc(ctx context.Context, docid uint64) {
	defer func(cost func() time.Duration) {
		log.Warnf("trackid:%v, cost: %.3f ms", ctx.Value("trackid"), float64(cost().Microseconds())/1000.0)
	}(tools.TimeCost())

	eg.delDoc(docid)
	log.Infof("trackid:%v, docid:%d", ctx.Value("trackid"), docid)
}

func DocIsDel(ctx context.Context, docid uint64) bool {
	defer func(cost func() time.Duration) {
		log.Warnf("trackid:%v, cost: %.3f ms", ctx.Value("trackid"), float64(cost().Microseconds())/1000.0)
	}(tools.TimeCost())

	deleted := eg.docIsDel(docid)
	log.Infof("trackid:%v, docid:%d, delete:%t", ctx.Value("trackid"), docid, deleted)
	return deleted
}
