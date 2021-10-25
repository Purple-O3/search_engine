package customnet

import (
	"context"
	"net/http"
	"search_engine/internal/service/engine"
	"search_engine/internal/service/objs"
	"search_engine/internal/util/idgenerator"
	"strconv"

	//"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type Net interface {
	StartNet(ip string, port string)
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
}

func newCustomHttp() *customHttp {
	return new(customHttp)
}

func (ch *customHttp) StartNet(ip string, port string) {
	/*router := gin.Default()
	equals
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())*/

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.POST("/add_doc", addDoc)
	router.POST("/add_doc_4test", addDoc4Test)
	router.GET("/del_doc", delDoc)
	router.GET("/doc_isdel", docIsDel)
	router.POST("/retrieve", retrieveDoc)
	//性能调试
	//pprof.Register(router)
	router.Run(ip + ":" + port)
}

func addDoc(ctx *gin.Context) {
	var respData RespData
	var docReq objs.Doc
	if err := ctx.BindJSON(&docReq); err != nil {
		respData.Code = -1
		respData.Message = err.Error()
		ctx.JSON(http.StatusOK, respData)
		return
	}
	trackid, err := strconv.ParseUint(ctx.GetHeader("X-Trackid"), 10, 64)
	if err != nil {
		trackid = uint64(idgenerator.Generate())
	}
	newCtx := context.WithValue(ctx, "trackid", trackid)
	docid := engine.AddDoc(newCtx, docReq)

	respData.Code = 0
	respData.Message = "ok"
	respData.Result.Docid = docid
	ctx.JSON(http.StatusOK, respData)
}

func addDoc4Test(ctx *gin.Context) {
	var respData RespData
	var docReq objs.Doc
	if err := ctx.BindJSON(&docReq); err != nil {
		respData.Code = -1
		respData.Message = err.Error()
		ctx.JSON(http.StatusOK, respData)
		return
	}
	trackid, err := strconv.ParseUint(ctx.GetHeader("X-Trackid"), 10, 64)
	if err != nil {
		trackid = uint64(idgenerator.Generate())
	}
	newCtx := context.WithValue(ctx, "trackid", trackid)
	var docid uint64 = 1388185366023311360
	docid = engine.AddDoc4Test(newCtx, docReq, docid)

	respData.Code = 0
	respData.Message = "ok"
	respData.Result.Docid = docid
	ctx.JSON(http.StatusOK, respData)
}

func delDoc(ctx *gin.Context) {
	var respData RespData
	docidString := ctx.Query("docid")
	docid, err := strconv.ParseUint(docidString, 10, 64)
	if err != nil {
		respData.Code = -1
		respData.Message = err.Error()
		ctx.JSON(http.StatusOK, respData)
		return
	}
	trackid, err := strconv.ParseUint(ctx.GetHeader("X-Trackid"), 10, 64)
	if err != nil {
		trackid = uint64(idgenerator.Generate())
	}
	newCtx := context.WithValue(ctx, "trackid", trackid)
	engine.DelDoc(newCtx, docid)

	respData.Code = 0
	respData.Message = "ok"
	ctx.JSON(http.StatusOK, respData)
}

func docIsDel(ctx *gin.Context) {
	var respData RespData
	docidString := ctx.Query("docid")
	docid, err := strconv.ParseUint(docidString, 10, 64)
	if err != nil {
		respData.Code = -1
		respData.Message = err.Error()
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
		respData.Code = 0
		respData.Message = "doc is delete"
	} else {
		respData.Code = 0
		respData.Message = "doc is not delete"
	}
	ctx.JSON(http.StatusOK, respData)
}

func retrieveDoc(ctx *gin.Context) {
	var respData RespData
	var rr RetreiveReq
	if err := ctx.BindJSON(&rr); err != nil {
		respData.Code = -1
		respData.Message = err.Error()
		ctx.JSON(http.StatusOK, respData)
		return
	}
	trackid, err := strconv.ParseUint(ctx.GetHeader("X-Trackid"), 10, 64)
	if err != nil {
		trackid = uint64(idgenerator.Generate())
	}
	newCtx := context.WithValue(ctx, "trackid", trackid)

	repl := engine.RetrieveDoc(newCtx, rr.RetreiveTerms)

	respData.Code = 0
	respData.Message = "ok"
	respData.Result.Repl = repl
	ctx.JSON(http.StatusOK, respData)
}
