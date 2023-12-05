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
		masterModels.Company{},
		masterModels.Client{},
		// common
		models.UserLog{},
		models.OperatingActivity{},
		models.Product{},
		models.Quotation{},
		models.ProductHistory{},
		models.Payment{},
		models.PaymentInstallment{},
		models.OperatingActivityProduct{},
		models.PurchaseOrder{},
		models.Document{},
		models.PurchaseOrderProduct{},
		models.DeliveryOrder{},
		models.Invoice{},
	); err != nil {
		panic(err)
	}
	fmt.Println("Database Migrated")

	database.DropUnusedColumns(masterModels.User{})
	database.DropUnusedColumns(masterModels.Company{})
	database.DropUnusedColumns(masterModels.Client{})

	database.DropUnusedColumns(models.UserLog{})
	database.DropUnusedColumns(models.OperatingActivity{})
	database.DropUnusedColumns(models.Product{})
	database.DropUnusedColumns(models.Quotation{})
	database.DropUnusedColumns(models.ProductHistory{})
	database.DropUnusedColumns(models.Payment{})
	database.DropUnusedColumns(models.PaymentInstallment{})
	database.DropUnusedColumns(models.OperatingActivityProduct{})
	database.DropUnusedColumns(models.PurchaseOrder{})
	database.DropUnusedColumns(models.Document{})
	database.DropUnusedColumns(models.PurchaseOrderProduct{})
	database.DropUnusedColumns(models.DeliveryOrder{})
	database.DropUnusedColumns(models.Invoice{})

	// foreign key queries
	//database.GetDatabaseConnection().Exec("ALTER TABLE companies ADD CONSTRAINT fk_created_by_user FOREIGN KEY (created_by) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE")
	//database.GetDatabaseConnection().Exec("ALTER TABLE companies ADD CONSTRAINT fk_updated_by_user FOREIGN KEY (updated_by) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE")

	//database.GetDatabaseConnection().Exec("ALTER TABLE clients ADD CONSTRAINT fk_created_by_user FOREIGN KEY (created_by) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE")
	//database.GetDatabaseConnection().Exec("ALTER TABLE clients ADD CONSTRAINT fk_updated_by_user FOREIGN KEY (updated_by) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE")

	//if err := database.GetDatabaseConnection().Exec("ALTER TABLE users ALTER COLUMN created_by DROP NOT NULL"); err != nil {
	//	panic(err)
	//}

	//if err := database.GetDatabaseConnection().Exec("ALTER TABLE users ALTER COLUMN updated_by DROP NOT NULL"); err != nil {
	//	panic(err)
	//}//

	//if err := database.GetDatabaseConnection().Exec("ALTER TABLE users ALTER COLUMN email SET NOT NULL"); err != nil {
	//	panic(err)
	//}
}
