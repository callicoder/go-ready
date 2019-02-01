package cmd

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/callicoder/go-ready/internal/config"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
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
	m, err := createMigrate(config)
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
	m, err := createMigrate(config)
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

func createMigrate(config *config.Config) (*migrate.Migrate, error) {
	if !strings.EqualFold(config.Database.DriverName, "postgres") {
		return nil, errors.New("Migratioin is supported only for postgres.")
	}

	db, err := sql.Open(config.Database.DriverName, config.Database.URL)
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(config.Migration.Path, "postgres", driver)
	if err != nil {
		return nil, err
	}

	return m, nil
}
