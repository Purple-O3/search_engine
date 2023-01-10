package controller

import (
	"fmt"
	"net/http"
	"runtime"
	"search_engine/internal/objs"
	enginepack "search_engine/internal/service/engine"
	"search_engine/internal/util/ginwrapper"
	"search_engine/internal/util/log"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func StartNet(config objs.ServerConfig, closeFunc func()) error {
	// handle options
	opts, err := setOpts(config)
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
	r.Use(recovery())
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

func recovery() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				respData := make(map[string]interface{})
				respData["code"] = -1
				respData["message"] = fmt.Sprintf("%v", err)
				ctx.JSON(http.StatusOK, respData)
				errBuf := make([]byte, 0, 1024)
				errBuf = errBuf[:runtime.Stack(errBuf, false)]
				log.Errorf("%s", string(errBuf))
				return
			}
		}()
		ctx.Next()
	}
}

func setOpts(config objs.ServerConfig) ([]ginwrapper.Option, error) {
	var opts []ginwrapper.Option
	if config.ReadTimeout != 0 {
		opts = append(opts, ginwrapper.WithReadTimeout(config.ReadTimeout))
	}

	if config.WriteTimeout != 0 {
		opts = append(opts, ginwrapper.WithWriteTimeout(config.WriteTimeout))
	}

	if config.IdleTimeout != 0 {
		opts = append(opts, ginwrapper.WithIdleTimeout(config.IdleTimeout))
	}
	opts = append(opts, ginwrapper.WithTLSConfig(config.Tls.Enable, config.Tls.CertFile, config.Tls.KeyFile))
	return opts, nil
}
