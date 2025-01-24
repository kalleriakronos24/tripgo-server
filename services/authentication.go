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
	"github.com/kalleriakronos24/khaimal-group/models"
	masterModels "github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/pkg/mail-service"
	"github.com/kalleriakronos24/khaimal-group/templates/email"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

func (module *module) AuthenticateUser(credentials dto.CredentialSignInDto) (token string, err error) {
	var user masterModels.Credentials
	if user, err = module.db.credentialModel.GetOneByEmail(credentials.Email); err != nil {
		return "", errors.New("email or password is incorrect")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return "", errors.New("email or password is incorrect")
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

// Internal / Dashboard user registration
func (module *module) RegisterUser(credentials *dto.UserSignup) (err error) {
	tx := database.GetDatabaseConnection().Begin()
	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost); err != nil {
		return errors.New("server error. please try again later")
	}

	var cred *masterModels.Credentials

	log.Printf("%v", credentials)
	if cred, err = module.db.credentialModel.InsertCredentials(masterModels.Credentials{
		Email:    credentials.Email,
		Password: string(hashedPassword),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	if err = module.db.userInternalModel.InsertInternal(masterModels.Internal{
		Name:          credentials.Name,
		Email:         cred.Email,
		Password:      string(hashedPassword),
		Phone:         "",
		CredentialsID: cred.ID,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  cred.Email,
		Subject: "Registration Success",
		Body:    email.ETRegisterSuccess(credentials.Name),
	}); err != nil {
		return errors.New("server error. please try again later")
	}
	tx.Commit()
	return
}

func (module *module) RegisterCustomer(credentials *dto.CustomerSignup) (err error) {
	tx := database.GetDatabaseConnection().Begin()
	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost); err != nil {
		return errors.New("server error. please try again later")
	}

	var cred *masterModels.Credentials

	if cred, err = module.db.credentialModel.InsertCredentials(masterModels.Credentials{
		Email:    credentials.Email,
		Password: string(hashedPassword),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	if err = module.db.userCustomerModel.InsertCustomer(masterModels.Customer{
		Name:          credentials.Name,
		Phone:         credentials.Phone,
		CredentialsID: cred.ID,
		RefferalCode:  credentials.RefferalCode,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  cred.Email,
		Subject: "Registration Success",
		Body:    email.ETRegisterSuccess(credentials.Name),
	}); err != nil {
		return errors.New("server error. please try again later")
	}
	tx.Commit()
	return
}

func (module *module) RegisterDriver(credentials *dto.DriverSignup) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost); err != nil {
		return errors.New("server error. please try again later")
	}

	var cred *masterModels.Credentials

	if cred, err = module.db.credentialModel.InsertCredentials(masterModels.Credentials{
		Email:    credentials.Email,
		Password: string(hashedPassword),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	if err = module.db.userDriverModel.InsertDriver(masterModels.Driver{
		Name:          credentials.Name,
		CredentialsID: cred.ID,
		DriverType:    "external",
		Phone:         credentials.Phone,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	var driver masterModels.Credentials
	if txError := tx.Model(&driver).
		Where("id = ?", cred.ID).
		Preload(clause.Associations).
		First(&driver); txError.Error != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}
	if err = module.db.balanceDriver.InsertBalanceDriver(models.BalanceDriver{
		Amount:   0,
		DriverID: driver.CredentialDriver.ID,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  cred.Email,
		Subject: "Registration Success",
		Body:    email.ETRegisterSuccess(credentials.Name),
	}); err != nil {
		return errors.New("server error. please try again later")
	}
	tx.Commit()
	return
}

func (module *module) RegisterDriverInternal(credentials *dto.DriverSignup) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost); err != nil {
		return errors.New("server error. please try again later")
	}

	var cred *masterModels.Credentials

	if cred, err = module.db.credentialModel.InsertCredentials(masterModels.Credentials{
		Email:    credentials.Email,
		Password: string(hashedPassword),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	if err = module.db.userDriverModel.InsertDriver(masterModels.Driver{
		Name:          credentials.Name,
		CredentialsID: cred.ID,
		DriverType:    credentials.DriverType,
		Phone:         credentials.Phone,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	var driver masterModels.Credentials
	if txError := tx.Model(&driver).
		Where("id = ?", cred.ID).
		Preload(clause.Associations).
		First(&driver); txError.Error != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}
	// todo
	// fix
	if err = module.db.balanceDriver.InsertBalanceDriver(models.BalanceDriver{
		Amount:   3000,
		DriverID: driver.CredentialDriver.ID,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  cred.Email,
		Subject: "Registration Success",
		Body:    email.ETRegisterSuccess(credentials.Name),
	}); err != nil {
		return errors.New("server error. please try again later")
	}
	tx.Commit()
	return
}

func (module *module) ResetPassword(p *dto.UniversalResetPassword) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost); err != nil {
		return errors.New("server error. please try again later")
	}

	if _, err = module.db.credentialModel.UpdateCredentials(masterModels.Credentials{
		Email:    p.Email,
		Password: string(hashedPassword),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to update password. try again")
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  p.Email,
		Subject: "Update Password Success",
		Body:    fmt.Sprintf("Your new password is: %v", p.Password),
	}); err != nil {
		return errors.New("server error. please try again later")
	}
	tx.Commit()
	return
}

func generateToken(user masterModels.Credentials) (string, error) {
	now := time.Now()
	expiry := time.Now().Add(constants.AuthenticationTimeout)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.JWTClaims{
		ID:        user.ID,
		ExpiresAt: expiry.Unix(),
		IssuedAt:  now.Unix(),
	})
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}
