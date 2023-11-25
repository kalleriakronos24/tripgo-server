package services

import (
	"gitlab.com/odma1/odma-be/config"
	"log"

	"github.com/google/uuid"
	"gitlab.com/odma1/odma-be/dto"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Handler HandlerFunc

type HandlerFunc interface {
	// AuthenticateUser ====== Authentication Handlers ====== //
	AuthenticateUser(p dto.UserLogin) (token string, err error)

	// RegisterUser ====== User Handlers ====== //
	RegisterUser(p dto.UserSignup) (err error)
	RetrieveUser(id uuid.UUID) (m masterModels.User, err error)
	UpdateUser(id uuid.UUID, p dto.UserUpdate) (err error)
}

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn      *gorm.DB
	userModel masterModels.UserModelAction
}

func InitializeServices() (err error) {
	// Initialize DB
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(config.AppConfig.DBUrl), &gorm.Config{})
	if err != nil {
		log.Println("[INIT] failed connecting to PostgresSQL")
		return
	}
	log.Println("[INIT] connected to PostgresSQL")
	///
	// Compose handler modules
	Handler = &module{
		db: &dbEntity{
			conn:      db,
			userModel: masterModels.NewUserAction(db),
		},
	}
	return
}
