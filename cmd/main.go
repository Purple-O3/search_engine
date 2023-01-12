package main

import (
	"search_engine/internal/controller"
	"search_engine/internal/objs"
	"search_engine/internal/service/engine"
	"search_engine/internal/util/log"
	"search_engine/internal/util/viperwrapper"

	_ "go.uber.org/automaxprocs"
)

func start() {
	configPath := "../configs/engine.yaml"
	var config objs.Config
	err := viperwrapper.DecodeConfig(configPath, &config)
	if err != nil {
		panic(err)
	}

	if config.Log.Type == "file" {
		log.InitLogger(config.Log)
	}
	engine.NewEngineWrap(config.Analyzer, config.DB, config.Bloomfilter)
	if err = controller.StartNet(config.Server, closeFunc); err != nil {
		panic(err)
	}
}

func closeFunc() {
	engine.CloseEngineWrap()
	log.CloseLogger()
}

func main() {
	start()
}
