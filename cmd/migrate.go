package cmd

import (
	"github.com/callicoder/go-ready/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

func newMigrateCommand() *cobra.Command {
	var migrateCmd = &cobra.Command{
		Use:          "migrate",
		Short:        "Run database migration",
		RunE:         migrateCmdF,
		SilenceUsage: true,
	}
	return migrateCmd
}

func newRollbackCommand() *cobra.Command {
	var rollbackCmd = &cobra.Command{
		Use:          "rollback",
		Short:        "Rollback database migration",
		RunE:         rollbackCmdF,
		SilenceUsage: true,
	}
	return rollbackCmd
}

func migrateCmdF(command *cobra.Command, args []string) error {
	configFileLocation, err := command.Flags().GetString("config")

	if err != nil {
		return err
	}

	config, err := config.Load(configFileLocation)

	if err != nil {
		return err
	}

	return runMigrations(config)
}

func rollbackCmdF(command *cobra.Command, args []string) error {
	configFileLocation, err := command.Flags().GetString("config")

	if err != nil {
		return err
	}

	config, err := config.Load(configFileLocation)
	if err != nil {
		return err
	}

	return rollbackMigration(config)
}

func runMigrations(config *config.Config) error {
	m, err := newMigrate(config)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}
		return err
	}

	return nil
}

func rollbackMigration(config *config.Config) error {
	m, err := newMigrate(config)
	if err != nil {
		return err
	}
	err = m.Steps(-1)
	if err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}
		return err
	}

	return nil
}

func newMigrate(config *config.Config) (*migrate.Migrate, error) {
	return migrate.New(config.Migration.Path, config.Database.URL)
}
