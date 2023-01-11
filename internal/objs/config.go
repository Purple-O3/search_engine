package objs

import (
	"search_engine/internal/util/bloomfilter"
	"search_engine/internal/util/ginwrapper"
	"search_engine/internal/util/log"
	"time"
)

type LogConfig = log.Config
type BloomfilterConfig = bloomfilter.Config
type ServerConfig = ginwrapper.Config

type Config struct {
	Server      ServerConfig
	Log         LogConfig
	Analyzer    AnalyzerConfig
	Bloomfilter BloomfilterConfig
	DB          DBConfig
}

type AnalyzerConfig struct {
	StopWordPath string
}

type DBConfig struct {
	Type     string
	Path     string
	Host     string
	Port     int
	Password string
	Index    int
	Timeout  time.Duration
}
