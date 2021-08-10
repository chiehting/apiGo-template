package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	_configFile      = "config/config.yml"
	_configEnvPrefix = "env"
)

// Application is define application
type Application struct {
	Name    string
	Version string
}

// Server is setting http server
type Server struct {
	RunMode        string
	HTTPPort       string
	MaxHeaderBytes int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

// Log is setting log
type Log struct {
	Level       string
	FilePath    string
	OmitTimeKey bool
}

// Database is setting database storage
type Database struct {
	Driver          string
	Host            string
	User            string
	Password        string
	DB              string
	Parameters      string
	MaxOpenConns    int
	MaxIdleConns    int
	MigrationFolder string
	LogLevel        int
}

// Cache is setting memory data storage
type Cache struct {
	Host     string
	DB       string
	Password string
}

type config struct {
	Application Application
	Server      Server
	Log         Log
	Database    Database
	Cache       Cache
}

var cfg *config

// Setup initialize the configuration instance
func init() {
	setConfigDefault()
	err := getConfig()
	if err != nil {
		panic(err)
	}
}

func setConfigDefault() {
	var null interface{}
	// application 預設值
	viper.SetDefault("application.name", "")

	// server 預設值
	viper.SetDefault("server.runmode", "release")
	viper.SetDefault("server.httpport", "80")
	viper.SetDefault("server.maxheaderbytes", 1<<20)
	viper.SetDefault("server.readtimeout", 60)
	viper.SetDefault("server.writetimeout", 60)

	// database 預設值
	viper.SetDefault("database.parameters", "?charset=utf8mb4&parseTime=True&loc=Local")
	viper.SetDefault("database.maxOpenConns", 10)
	viper.SetDefault("database.maxIdleConns", 10)
	viper.SetDefault("database.migrationFolder", "./sql")
	viper.SetDefault("database.loglevel", 1)

	// cache 預設值
	viper.SetDefault("cache.db", "0")

	// log 預設值
	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.filePath", null)
	viper.SetDefault("log.omittimekey", false)
}

func getConfig() (err error) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(_configEnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigFile(_configFile)
	if err := viper.ReadInConfig(); err == nil {
		viper.Unmarshal(&cfg)
	}

	cfg.prepareConfig()

	return
}

// prepareConfig 調整設定檔
func (cfg *config) prepareConfig() {
	cfg.Log.Level = strings.ToLower(cfg.Log.Level)
	if cfg.Log.Level == "debug" {
		cfg.Server.RunMode = cfg.Log.Level
		cfg.Database.LogLevel = 4
	}

	cfg.Server.HTTPPort = ":" + cfg.Server.HTTPPort
	cfg.Server.ReadTimeout = cfg.Server.ReadTimeout * time.Second
	cfg.Server.WriteTimeout = cfg.Server.WriteTimeout * time.Second

	if len(cfg.Application.Name) == 0 {
		panic("Please set the application name to config file or environment variable!")
	}
}

// GetApplication 取得服務資訊
func GetApplication() *Application {
	return &cfg.Application
}

// GetServer 取得伺服器資訊
func GetServer() *Server {
	return &cfg.Server
}

// GetDatabase 取得資料庫連線資訊
func GetDatabase() *Database {
	return &cfg.Database
}

// GetCache 取得快取連線資訊
func GetCache() *Cache {
	return &cfg.Cache
}

// GetLog 取得日誌資訊
func GetLog() *Log {
	return &cfg.Log
}
