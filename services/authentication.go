package services

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"gitlab.com/odma1/odma-be/config"
	"gitlab.com/odma1/odma-be/constants"
	"gitlab.com/odma1/odma-be/dto"
	masterModels "gitlab.com/odma1/odma-be/models/master"
	"golang.org/x/crypto/bcrypt"
)

func (module *module) AuthenticateUser(credentials dto.UserLogin) (token string, err error) {
	var user masterModels.User
	if user, err = module.db.userModel.GetOneByUsername(credentials.Username); err != nil {
		return "", errors.New("incorrect credentials")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return "", errors.New("incorrect credentials")
	}

	return generateToken(user)
}

func (module *module) RegisterUser(credentials dto.UserSignup) (err error) {
	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost); err != nil {
		return errors.New("failed hashing password")
	}
	if _, err = module.db.userModel.InsertUser(masterModels.User{
		Username: credentials.Username,
		Email:    credentials.Email,
		Password: string(hashedPassword),
		Bio:      credentials.Bio,
	}); err != nil {
		log.Print(err)
		return fmt.Errorf("error inserting user. %v", err)
	}
	return
}

func generateToken(user masterModels.User) (string, error) {
	now := time.Now()
	expiry := time.Now().Add(constants.AuthenticationTimeout)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.JWTClaims{
		ID:        user.ID,
		ExpiresAt: expiry.Unix(),
		IssuedAt:  now.Unix(),
	})
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}
