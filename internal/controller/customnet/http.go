package customnet

import (
	"context"
	"net/http"
	"search_engine/internal/service/engine"
	"search_engine/internal/service/objs"
	"search_engine/internal/util/idgenerator"
	"search_engine/internal/util/log"
	"strconv"

	//"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type Net interface {
	StartNet(ip string, port int)
	Shutdown()
}

func NetFactory(netType string) Net {
	switch netType {
	case "http":
		return newCustomHttp()
	//TODO case "rpc":
	default:
		return newCustomHttp()
	}
}

type customHttp struct {
	svr *http.Server
}

func newCustomHttp() *customHttp {
	return new(customHttp)
}

func (ch *customHttp) StartNet(ip string, port int) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.POST("/add_doc", addDoc)
	router.GET("/del_doc", delDoc)
	router.GET("/doc_isdel", docIsDel)
	router.POST("/retrieve", retrieveDoc)
	//性能调试
	//pprof.Register(router)

	srv := &http.Server{
		Addr:    ip + ":" + strconv.Itoa(port),
		Handler: router,
	}
	ch.svr = srv
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("%v", err)
		}
	}()
}

func (ch *customHttp) Shutdown() {
	ch.svr.Shutdown(context.Background())
}

func addDoc(ctx *gin.Context) {
	respData := make(map[string]interface{})
	var docReq objs.Doc
	if err := ctx.BindJSON(&docReq); err != nil {
		respData["code"] = -1
		respData["message"] = err.Error()
		ctx.JSON(http.StatusOK, respData)
		return
	}
	trackid, err := strconv.ParseUint(ctx.GetHeader("X-Trackid"), 10, 64)
	if err != nil {
		trackid = uint64(idgenerator.Generate())
	}
	newCtx := context.WithValue(ctx, "trackid", trackid)

	docid := engine.AddDoc(newCtx, docReq)

	respData["code"] = 0
	respData["message"] = "ok"
	respData["docid"] = docid
	ctx.JSON(http.StatusOK, respData)
}

func delDoc(ctx *gin.Context) {
	respData := make(map[string]interface{})
	docidString := ctx.Query("docid")
	docid, err := strconv.ParseUint(docidString, 10, 64)
	if err != nil {
		respData["code"] = -1
		respData["message"] = err.Error()
		ctx.JSON(http.StatusOK, respData)
		return
	}
	trackid, err := strconv.ParseUint(ctx.GetHeader("X-Trackid"), 10, 64)
	if err != nil {
		trackid = uint64(idgenerator.Generate())
	}
	newCtx := context.WithValue(ctx, "trackid", trackid)

	engine.DelDoc(newCtx, docid)

	respData["code"] = 0
	respData["message"] = "ok"
	ctx.JSON(http.StatusOK, respData)
}

func docIsDel(ctx *gin.Context) {
	respData := make(map[string]interface{})
	docidString := ctx.Query("docid")
	docid, err := strconv.ParseUint(docidString, 10, 64)
	if err != nil {
		respData["code"] = -1
		respData["message"] = err.Error()
		ctx.JSON(http.StatusOK, respData)
		return
	}
	trackid, err := strconv.ParseUint(ctx.GetHeader("X-Trackid"), 10, 64)
	if err != nil {
		trackid = uint64(idgenerator.Generate())
	}
	newCtx := context.WithValue(ctx, "trackid", trackid)

	del := engine.DocIsDel(newCtx, docid)

	if del {
		respData["code"] = 0
		respData["message"] = "doc is delete"
	} else {
		respData["code"] = 0
		respData["message"] = "doc is not delete"
	}
	ctx.JSON(http.StatusOK, respData)
}

func retrieveDoc(ctx *gin.Context) {
	respData := make(map[string]interface{})
	var rr objs.RetreiveReq
	if err := ctx.BindJSON(&rr); err != nil {
		respData["code"] = -1
		respData["message"] = err.Error()
		ctx.JSON(http.StatusOK, respData)
		return
	}
	trackid, err := strconv.ParseUint(ctx.GetHeader("X-Trackid"), 10, 64)
	if err != nil {
		trackid = uint64(idgenerator.Generate())
	}
	newCtx := context.WithValue(ctx, "trackid", trackid)

	repl := engine.RetrieveDoc(newCtx, rr.RetreiveTerms)
	replLen := len(repl)
	if rr.Offset == 0 && rr.Limit == 0 {
		rr.Limit = 10
	}
	end := rr.Offset + rr.Limit
	if replLen >= end {
		repl = repl[:end]
	}

	respData["code"] = 0
	respData["message"] = "ok"
	respData["count"] = replLen
	respData["result"] = repl
	ctx.JSON(http.StatusOK, respData)
}
