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
	); err != nil {
		panic(err)
	}
	fmt.Println("Database Migrated")

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
