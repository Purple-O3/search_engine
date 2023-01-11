package main

import (
	"errors"
	"path/filepath"
	"search_engine/internal/controller"
	"search_engine/internal/objs"
	"search_engine/internal/service/engine"
	"search_engine/internal/util/log"
	"strings"

	"github.com/spf13/viper"
	_ "go.uber.org/automaxprocs"
)

func start() {
	configPath := "../configs/engine.yaml"
	fileName := filepath.Base(configPath)
	fileNames := strings.Split(fileName, ".")
	if len(fileNames) != 2 {
		panic(errors.New("fileNames len not equal 2"))
	}
	var config objs.Config
	vp := viper.New()
	vp.AddConfigPath(filepath.Dir(configPath))
	vp.SetConfigName(fileNames[0])
	vp.SetConfigType(fileNames[1])
	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = vp.Unmarshal(&config)
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
