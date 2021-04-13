package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	configFile      = "config/config.yml"
	configEnvPrefix = "env"
)

type server struct {
	RunMode        string
	HTTPPort       string
	MaxHeaderBytes int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

type log struct {
	Level       string
	FilePath    string
	OmitTimeKey bool
}

type database struct {
	Driver          string
	User            string
	Password        string
	Host            string
	Name            string
	MaxOpenConns    int
	MaxIdleConns    int
	MigrationFolder string
}

type config struct {
	Server   server
	Log      log
	Database database
}

// Setup initialize the configuration instance
func init() {
	setConfigDefault()
	err := getConfig()
	if err != nil {
		panic(err)
	}
}

var cfg *config

func setConfigDefault() {
	var null interface{}
	// server 預設值
	viper.SetDefault("server.runmode", "debug")
	viper.SetDefault("server.httpport", "80")
	viper.SetDefault("server.maxheaderbytes", 1<<20)
	viper.SetDefault("server.readtimeout", 60)
	viper.SetDefault("server.writetimeout", 60)

	// database 預設值
	viper.SetDefault("database.maxOpenConns", 10)
	viper.SetDefault("database.maxIdleConns", 10)
	viper.SetDefault("database.migrationFolder", "./sql")

	// log 預設值
	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.filePath", null)
	viper.SetDefault("log.omittimekey", false)
}

func getConfig() (err error) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(configEnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigFile(configFile)
	err = viper.ReadInConfig()

	if err == nil {
		err = viper.Unmarshal(&cfg)
	}

	cfg.prepareConfig()
	return
}

// prepareConfig 準備設定檔
func (cfg *config) prepareConfig() {
	cfg.Server.RunMode = "release"
	if cfg.Log.Level == "debug" {
		cfg.Server.RunMode = cfg.Log.Level
	}
	cfg.Server.HTTPPort = ":" + cfg.Server.HTTPPort
	cfg.Server.ReadTimeout = cfg.Server.ReadTimeout * time.Second
	cfg.Server.WriteTimeout = cfg.Server.WriteTimeout * time.Second
}

// GetDatabaseDns 取得日誌資訊
func GetLogConfig() *log {
	return &cfg.Log
}

// GetDatabaseConfig 取得資料庫連線資訊
func GetDatabaseConfig() *database {
	return &cfg.Database
}

// GetServerConfig 取得伺服器資訊
func GetServerConfig() *server {
	return &cfg.Server
}
