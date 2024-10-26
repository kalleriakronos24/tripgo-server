package services

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/config"
	"github.com/kalleriakronos24/khaimal-group/constants"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/dto"
	masterModels "github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/pkg/mail-service"
	"golang.org/x/crypto/bcrypt"
)

func (module *module) AuthenticateUser(credentials dto.CredentialSignInDto) (token string, err error) {
	var user masterModels.Credentials
	if user, err = module.db.credentialModel.GetOneByEmail(credentials.Email); err != nil {
		return "", errors.New("incorrect credentials")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return "", errors.New("incorrect credentials")
	}
	return generateToken(user)
}

func (module *module) UpdateDeviceToken(token string, credentialId uuid.UUID) (err error) {
	if token != "" {
		if err = module.db.credentialModel.UpdateDeviceToken(credentialId, masterModels.Credentials{
			DeviceToken: token,
		}); err != nil {
			return fmt.Errorf("failed update device token. %v", err)
		}
	}
	return err
}

func (module *module) RegisterUser(credentials *dto.UserSignup) (err error) {
	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost); err != nil {
		return errors.New("failed hashing password")
	}

	if err = module.db.userModel.InsertUser(masterModels.User{
		Name:      credentials.Name,
		Email:     credentials.Email,
		Password:  string(hashedPassword),
		CreatedBy: credentials.CreatedBy,
		UpdatedBy: credentials.CreatedBy,
		Role:      "user",
	}); err != nil {
		return fmt.Errorf("error inserting user. %v", err)
	}

	// mailPayload := &mail.TSendMail{}
	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "notification@wadahgo.com",
		MailTo:  "credentials.Email",
		Subject: "WadahGo - Registration Success",
		Body: `<html><body>
		<p>Thank You for registrering</p>
		</body></html>`,
	}); err != nil {
		return errors.New(err.Error())
	}
	return
}

func (module *module) RegisterCustomer(credentials *dto.CustomerSignup) (err error) {
	tx := database.GetDatabaseConnection().Begin()
	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost); err != nil {
		return errors.New("failed hashing password")
	}

	var cred *masterModels.Credentials

	if cred, err = module.db.credentialModel.InsertCredentials(masterModels.Credentials{
		Email:    credentials.Email,
		Password: string(hashedPassword),
	}, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting credential. %v", err)
	}

	if err = module.db.userCustomerModel.InsertCustomer(masterModels.Customer{
		Name:          credentials.Name,
		Phone:         credentials.Phone,
		CredentialsID: cred.ID,
	}, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting customer. %v", err)
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  cred.Email,
		Subject: "Registration Success",
		Body: `<html><body>
		<p>Thank You for registrering</p>
		</body></html>`,
	}); err != nil {
		return errors.New(err.Error())
	}
	tx.Commit()
	return
}

func (module *module) RegisterDriver(credentials *dto.DriverSignup) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost); err != nil {
		return errors.New("failed hashing password")
	}

	var cred *masterModels.Credentials

	if cred, err = module.db.credentialModel.InsertCredentials(masterModels.Credentials{
		Email:    credentials.Email,
		Password: string(hashedPassword),
	}, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting credential. %v", err)
	}

	if err = module.db.userDriverModel.InsertDriver(masterModels.Driver{
		Name:          credentials.Name,
		CredentialsID: cred.ID,
		DriverType:    credentials.DriverType,
		Phone:         credentials.Phone,
	}, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting driver. %v", err)
	}
	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  cred.Email,
		Subject: "Registration Success",
		Body: `<html><body>
		<p>Thank You for registrering</p>
		</body></html>`,
	}); err != nil {
		return errors.New(err.Error())
	}
	tx.Commit()
	return
}

func generateToken(user masterModels.Credentials) (string, error) {

	log.Printf("%v", user)
	now := time.Now()
	expiry := time.Now().Add(constants.AuthenticationTimeout)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.JWTClaims{
		ID:        user.ID,
		ExpiresAt: expiry.Unix(),
		IssuedAt:  now.Unix(),
	})
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}
