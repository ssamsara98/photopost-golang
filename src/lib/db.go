package lib

import (
	"fmt"
	"go-photopost/src/entities"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(logger *zap.Logger, env *Env) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		env.DBHost,
		env.DBUsername,
		env.DBPassword,
		env.DBName,
		env.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Sugar().Infoln(err.Error())
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(
		&entities.User{},
		&entities.Post{},
		&entities.PostPhoto{},
		&entities.PostToPhoto{},
	)

	return db
}
