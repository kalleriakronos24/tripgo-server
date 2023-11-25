package migrations

import (
	"fmt"

	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/models"
	masterModels "gitlab.com/odma1/odma-be/models/master"
)

func Migrate() {
	// auto migration all models
	if err := database.GetDatabaseConnection().AutoMigrate(
		// master
		masterModels.User{},
		// common
		models.UserLog{},
	); err != nil {
		panic(err)
	}
	fmt.Println("Database Migrated")
}
