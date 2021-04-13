package database

import (
	"fmt"

	"github.com/chiehting/go-template/pkg/config"

	command "github.com/chiehting/db-migrate/command"
	migrate "github.com/chiehting/db-migrate/migrate"
)

// RunMigration 執行資料庫建立
func RunMigration() error {
	cfg := config.GetDatabaseConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Name)

	env := &command.Environment{
		Dialect:    cfg.Driver,
		DataSource: dsn,
		Dir:        cfg.MigrationFolder,
	}
	command.SetEnvironment(env)
	migrate.SetIgnoreUnknown(true)
	Upcommand := command.UpCommand{}

	return Upcommand.RunProcess([]string{})
}
