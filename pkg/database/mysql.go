package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/chiehting/go-template/pkg/config"
	"github.com/chiehting/go-template/pkg/log"
	_ "github.com/go-sql-driver/mysql"
)

func ConnDB() *sql.DB {
	cfg := config.GetDatabaseConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Name)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		log.Panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	err = db.Ping()
	log.Info(db.Stats())
	if err != nil {
		log.Panic(db.Ping())
	}

	return db
}
