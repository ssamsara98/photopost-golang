package infrastructure

import (
	"fmt"
	"go-clean-arch/lib"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database modal
type Database struct {
	*gorm.DB
	logger lib.Logger
}

// NewDatabase creates a new database instance
func NewDatabase(
	env *lib.Env,
	log lib.Logger,
) Database {
	url := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		env.DBUsername,
		env.DBPassword,
		env.DBHost,
		env.DBPort,
		env.DBName,
	)

	log.Info("opening db connection")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: log.GetGormLogger(),
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Info("Url: ", url)
		log.Panic(err)
	}

	log.Info("database connection established")

	database := Database{
		db,
		log,
	}
	log.Info("currentDatabase:", db.Migrator().CurrentDatabase())

	if err := RunMigration(env, log, database); err != nil {
		log.Info("migration failed.")
		log.Panic(err)
	}

	return database
}

// WithTrx delegate transaction from user repository
func (d Database) WithTrx(trxHandle *gorm.DB) Database {
	if trxHandle != nil {
		d.logger.Debug("using WithTrx as trxHandle is not nil")
		d.DB = trxHandle
	}
	return d
}
