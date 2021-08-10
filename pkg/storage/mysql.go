package storage

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/chiehting/apiGo-template/pkg/config"
)

// MySQL is connection to database
var MySQL = mysqlConnection()

func mysqlConnection() *gorm.DB {
	cfg := config.GetDatabase()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.DB+cfg.Parameters,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.LogLevel(cfg.LogLevel), // Log level
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic(err)
	}

	// Setting MySQL connection pool parameters
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	return db
}
