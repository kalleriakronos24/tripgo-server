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
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  cred.Email,
		Subject: "Registration Success",
		Body: `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome to Mailmeteor</title>
    <!--[if mso]><style type="text/css">body, table, td, a { font-family: Arial, Helvetica, sans-serif !important; }</style><![endif]-->
</head>
<body style="font-family: Helvetica, Arial, sans-serif; margin: 0px; padding: 0px; background-color: #ffffff;">
    <table role="presentation"
        style="width: 100%; border-collapse: collapse; border: 0px; border-spacing: 0px; font-family: Arial, Helvetica, sans-serif; background-color: rgb(239, 239, 239);">
        <tbody>
            <tr>
                <td align="center" style="padding: 1rem 2rem; vertical-align: top; width: 100%;">
                    <table role="presentation"
                        style="max-width: 600px; border-collapse: collapse; border: 0px; border-spacing: 0px; text-align: left;">
                        <tbody>
                            <tr>
                                <td style="padding: 40px 0px 0px;">
                                    <div style="padding: 20px; background-color: rgb(255, 255, 255);">
                                        <div style="color: rgb(0, 0, 0); text-align: left;">
                                            <h1 style="margin: 1rem 0">Welcome to WadahGo</h1>
                                            <p style="padding-bottom: 16px">Hello ` + credentials.Name + `,</p>
                                            <p style="padding-bottom: 16px">Thank you for signing up to WadahGo.
                                                We're really happy to have you onboard! Click the
                                                link below to login to your account:</p>
                                            <p style="padding-bottom: 16px"><a href="https://wadahgo.com"
                                                    target="_blank"
                                                    style="padding: 12px 24px; border-radius: 4px; color: #FFF; background: #2B52F5;display: inline-block;margin: 0.5rem 0;">Login
                                                    to your account</a></p>
                                            <p style="padding-bottom: 16px">Best regards,</p><span
                                                style="color: #999">WadahGo Limousine</span></p>
                                        </div>
                                    </div>
                                    <div style="padding-top: 20px; color: rgb(153, 153, 153); text-align: center;">
                                        <p style="padding-bottom: 6px;">©Wadahgo 2024, Wadahgo™ is trademark of Wadah
                                            Hub
                                            All rights reserved.</p>
                                        <p>42-01, Jalan Sri Perkasa 1/3, 81200 Johor Bahru</p>
                                    </div>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </td>
            </tr>
        </tbody>
    </table>
</body>
</html>`,
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
		DriverType:    credentials.DriverType,
		Phone:         credentials.Phone,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}
	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  cred.Email,
		Subject: "Registration Success",
		Body: `<html><body>
		<p>Thank You for registrering</p>
		</body></html>`,
	}); err != nil {
		return errors.New("server error. please try again later")
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
