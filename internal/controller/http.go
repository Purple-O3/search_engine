package controller

import (
	"search_engine/internal/objs"
	enginepack "search_engine/internal/service/engine"
	"search_engine/internal/util/ginwrapper"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func StartNet(config objs.ServerConfig, closeFunc func()) error {
	opts, err := ginwrapper.SetOpts(config)
	if err != nil {
		return err
	}
	return ginwrapper.GinServer(config.IP, config.Port, router(config), closeFunc, opts...)
}

func router(config objs.ServerConfig) *gin.Engine {
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(ginwrapper.Recovery())
	apiGroup := r.Group("/api/v1")
	{
		apiGroup.POST("/add_doc", enginepack.AddDoc)
		apiGroup.GET("/del_doc", enginepack.DelDoc)
		apiGroup.GET("/doc_isdel", enginepack.DocIsDel)
		apiGroup.POST("/retrieve_doc", enginepack.RetrieveDoc)
	}
	if config.Debug {
		pprof.Register(r)
	}
	return r
}
