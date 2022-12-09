package main

import (
	"errors"
	"os"
	"os/signal"
	"path/filepath"
	"search_engine/internal/controller/customnet"
	"search_engine/internal/service/engine"
	"search_engine/internal/util/log"
	"strings"
	"syscall"

	"github.com/spf13/viper"
	_ "go.uber.org/automaxprocs"
)

var cn customnet.Net

func init() {
	var filePath string
	if len(os.Args) < 2 || os.Args[1] == "7788" {
		filePath = "../configs/engine.toml"
	} else if os.Args[1] == "7799" {
		filePath = "../configs/engine2.toml"
	}
	fileName := filepath.Base(filePath)
	fileNames := strings.Split(fileName, ".")
	if len(fileNames) != 2 {
		panic(errors.New("fileNames len not equal 2"))
	}

	vp := viper.New()
	vp.SetConfigName(fileNames[0])
	vp.SetConfigType(fileNames[1])
	vp.AddConfigPath(filepath.Dir(filePath))
	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}

	logLevel := vp.GetString("log.level")
	logFilePath := vp.GetString("log.file_path")
	logMaxSize := vp.GetInt("log.max_size")
	logMaxBackups := vp.GetInt("log.max_backups")
	logMaxAge := vp.GetInt("log.max_age")
	logCompress := vp.GetBool("log.compress")
	log.InitLogger(logLevel, logFilePath, logMaxSize, logMaxBackups, logMaxAge, logCompress)

	analyzerStopWordPath := vp.GetString("analyzer.stop_word_path")
	dbPath := vp.GetString("db.path")
	dbHost := vp.GetString("db.host")
	dbPort := vp.GetString("db.port")
	dbPassword := vp.GetString("db.password")
	dbIndex := vp.GetInt("db.index")
	dbTimeout := vp.GetInt("db.timeout")
	bloomfilterMiscalRate := vp.GetFloat64("bloomfilter.miscal_rate")
	bloomfilterAddSize := vp.GetUint64("bloomfilter.add_size")
	bloomfilterStorePath := vp.GetString("bloomfilter.store_path")
	engine.NewEg(analyzerStopWordPath, dbPath, dbHost, dbPort, dbPassword, dbIndex, dbTimeout, bloomfilterMiscalRate, bloomfilterAddSize, bloomfilterStorePath)

	ip := vp.GetString("server.ip")
	port := vp.GetInt("server.port")
	cn = customnet.NetFactory("http")
	cn.StartNet(ip, port)
	log.Infof("server start!!!")
}

func listenSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c
}

func closeServer() {
	engine.CloseEg()
	log.CloseLogger()
	cn.Shutdown()
}

func main() {
	listenSignal()
	closeServer()
}
