package objs

import (
	"search_engine/internal/util/bloomfilter"
	"search_engine/internal/util/log"
	"time"
)

type LogConfig = log.Config
type BloomfilterConfig = bloomfilter.Config

type Config struct {
	Server      ServerConfig
	Log         LogConfig
	Analyzer    AnalyzerConfig
	Bloomfilter BloomfilterConfig
	DB          DBConfig
}

type ServerConfig struct {
	Name         string
	IP           string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	Debug        bool
	Tls          TLS
}

type TLS struct {
	Enable   bool
	CertFile string
	KeyFile  string
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
