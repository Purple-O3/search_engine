package main

import (
	"os"
	"os/signal"
	"search_engine/internal/controller/customnet"
	"search_engine/internal/service/engine"
	"search_engine/internal/util/log"
	"syscall"

	"github.com/spf13/viper"
	_ "go.uber.org/automaxprocs"
)

var cn customnet.Net

func init() {
	if len(os.Args) < 2 || os.Args[1] == "7788" {
		viper.SetConfigName("engine")
	} else if os.Args[1] == "7799" {
		viper.SetConfigName("engine2")
	}
	viper.SetConfigType("toml")       // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("../configs") // 查找配置文件所在的路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	logLevel := viper.GetString("log.level")
	logFilePath := viper.GetString("log.file_path")
	logMaxSize := viper.GetInt("log.max_size")
	logMaxBackups := viper.GetInt("log.max_backups")
	logMaxAge := viper.GetInt("log.max_age")
	logCompress := viper.GetBool("log.compress")
	log.InitLogger(logLevel, logFilePath, logMaxSize, logMaxBackups, logMaxAge, logCompress)

	analyzerStopWordPath := viper.GetString("analyzer.stop_word_path")
	dbPath := viper.GetString("db.path")
	dbHost := viper.GetString("db.host")
	dbPort := viper.GetString("db.port")
	dbPassword := viper.GetString("db.password")
	dbIndex := viper.GetInt("db.index")
	dbTimeout := viper.GetInt("db.timeout")
	bloomfilterMiscalRate := viper.GetFloat64("bloomfilter.miscal_rate")
	bloomfilterAddSize := viper.GetUint64("bloomfilter.add_size")
	bloomfilterStorePath := viper.GetString("bloomfilter.store_path")
	engine.NewEg(analyzerStopWordPath, dbPath, dbHost, dbPort, dbPassword, dbIndex, dbTimeout, bloomfilterMiscalRate, bloomfilterAddSize, bloomfilterStorePath)

	ip := viper.GetString("server.ip")
	port := viper.GetString("server.port")
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
