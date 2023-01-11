package enginepack

import (
	"errors"
	"search_engine/internal/objs"
	"search_engine/internal/util/ginwrapper"
	"search_engine/internal/util/idgenerator"
	"search_engine/internal/util/log"
	"search_engine/internal/util/tools"
	"time"

	"github.com/gin-gonic/gin"
)

type engineWrapper struct {
	*engine
	ginwrapper.Base
}

var egw *engineWrapper

func NewEngineWrap(analyzerConfig objs.AnalyzerConfig, dbConfig objs.DBConfig, bloomfilterConfig objs.BloomfilterConfig) {
	egw = new(engineWrapper)
	egw.engine = newEngine(analyzerConfig, dbConfig, bloomfilterConfig)
}

func CloseEngineWrap() {
	egw.engine.close()
}

func RetrieveDoc(ctx *gin.Context) {
	var req objs.RetreiveReq
	if err := egw.BindAndValidate(ctx, &req); err != nil {
		egw.ErrMsg(ctx, err)
		return
	}
	if req.Trackid == 0 {
		req.Trackid = uint64(idgenerator.Generate())
	}
	defer func(cost func() time.Duration) {
		log.Warnf("trackid:%v, cost: %.3f ms", req.Trackid, float64(cost().Microseconds())/1000.0)
	}(tools.TimeCost())

	repl := egw.retrieveDoc(req.RetreiveTerms, req.Trackid)
	replLen := len(repl)
	if req.Offset == 0 && req.Limit == 0 {
		req.Limit = 10
	}
	end := req.Offset + req.Limit
	if replLen >= end {
		repl = repl[:end]
	}
	log.Infof("trackid:%v, RetreiveTerm:%v, repl:%v", req.Trackid, req.RetreiveTerms, repl)
	egw.SucMsg(ctx, objs.RetreiveDocResp{Count: replLen, Result: repl})
}

func AddDoc(ctx *gin.Context) {
	var req objs.DocReq
	if err := egw.BindAndValidate(ctx, &req); err != nil {
		egw.ErrMsg(ctx, err)
		return
	}
	if req.Trackid == 0 {
		req.Trackid = uint64(idgenerator.Generate())
	}
	docid := uint64(idgenerator.Generate())
	defer func(cost func() time.Duration) {
		log.Warnf("trackid:%v, cost: %.3f ms", req.Trackid, float64(cost().Microseconds())/1000.0)
	}(tools.TimeCost())

	egw.addDoc(req.Doc, docid, req.Trackid)
	log.Infof("trackid:%v, docid:%d, doc:%v", req.Trackid, docid, req.Doc)
	egw.SucMsg(ctx, objs.AddDocResp{Docid: docid})
}

func DelDoc(ctx *gin.Context) {
	var req objs.DocidReq
	if err := egw.BindAndValidate(ctx, &req); err != nil {
		egw.ErrMsg(ctx, err)
		return
	}
	if req.Trackid == 0 {
		req.Trackid = uint64(idgenerator.Generate())
	}
	defer func(cost func() time.Duration) {
		log.Warnf("trackid:%v, cost: %.3f ms", req.Trackid, float64(cost().Microseconds())/1000.0)
	}(tools.TimeCost())

	egw.delDoc(req.Docid)
	log.Infof("trackid:%v, docid:%d", req.Trackid, req.Docid)
	egw.SucMsg(ctx)
}

func DocIsDel(ctx *gin.Context) {
	var req objs.DocidReq
	if err := egw.BindAndValidate(ctx, &req); err != nil {
		egw.ErrMsg(ctx, err)
		return
	}
	if req.Trackid == 0 {
		req.Trackid = uint64(idgenerator.Generate())
	}
	defer func(cost func() time.Duration) {
		log.Warnf("trackid:%v, cost: %.3f ms", req.Trackid, float64(cost().Microseconds())/1000.0)
	}(tools.TimeCost())

	deleted := egw.docIsDel(req.Docid)
	log.Infof("trackid:%v, docid:%d, delete:%t", req.Trackid, req.Docid, deleted)
	if deleted {
		egw.ErrMsg(ctx, errors.New("doc is deleted"))
	} else {
		egw.ErrMsg(ctx, errors.New("doc is not deleted"))
	}
}
