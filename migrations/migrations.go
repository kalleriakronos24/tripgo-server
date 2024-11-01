package migrations

import (
	"fmt"

	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/models"
	masterModels "github.com/kalleriakronos24/khaimal-group/models/master"
)

func Migrate() {
	// auto migration all models
	if err := database.GetDatabaseConnection().AutoMigrate(
		// master
		masterModels.User{},
		masterModels.Credentials{},
		masterModels.Customer{},
		masterModels.Driver{},
		masterModels.Internal{},
		masterModels.Company{},
		masterModels.CarManagement{},
		masterModels.CarModel{},
		models.BookingTransfer{},
		models.BookingTransferAssigned{},
		models.BalanceDriver{},
		models.BookingTransferRating{},
	); err != nil {
		panic(err)
	}
	fmt.Println("Database Migrated")
	database.DropUnusedColumns(masterModels.User{})
	database.DropUnusedColumns(masterModels.Credentials{})
	database.DropUnusedColumns(masterModels.Customer{})
	database.DropUnusedColumns(masterModels.Driver{})
	database.DropUnusedColumns(masterModels.Internal{})
	database.DropUnusedColumns(masterModels.Company{})
	database.DropUnusedColumns(masterModels.CarManagement{})
	database.DropUnusedColumns(masterModels.CarModel{})
	database.DropUnusedColumns(models.BookingTransfer{})
	database.DropUnusedColumns(models.BookingTransferAssigned{})
	database.DropUnusedColumns(models.BalanceDriver{})
	database.DropUnusedColumns(models.BookingTransferRating{})
}
