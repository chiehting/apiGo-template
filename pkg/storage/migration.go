package storage

import (
	"fmt"

	"github.com/chiehting/apiGo-template/pkg/config"

	command "github.com/chiehting/db-migrate/command"
)

// Migration 執行資料庫建立
func Migration() {
	cfg := config.GetDatabase()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.DB+cfg.Parameters,
	)

	env := &command.Environment{
		Dialect:    cfg.Driver,
		DataSource: dsn,
		Dir:        cfg.MigrationFolder,
	}

	command.SetEnvironment(env)
	command.SetIgnoreUnknown(true)
	Upcommand := command.UpCommand{}

	if err := Upcommand.RunProcess([]string{}); err != nil {
		panic(err)
	}
}
