package database

import (
	"log"

	"gitlab.com/odma1/odma-be/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseConnection() *gorm.DB {
	var db *gorm.DB
	db, err := gorm.Open(postgres.Open(config.AppConfig.DBUrl), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		log.Println("[INIT] failed connecting to PostgresSQL")
		return nil
	}
	//
	return db
}
