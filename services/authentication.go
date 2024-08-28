package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/kalleriakronos24/khaimal-group/config"
	"github.com/kalleriakronos24/khaimal-group/constants"
	"github.com/kalleriakronos24/khaimal-group/dto"
	masterModels "github.com/kalleriakronos24/khaimal-group/models/master"
	"golang.org/x/crypto/bcrypt"
)

func (module *module) AuthenticateUser(credentials dto.UserLogin) (token string, err error) {
	var user masterModels.User
	if user, err = module.db.userModel.GetOneByEmail(credentials.Email); err != nil {
		return "", errors.New("incorrect credentials")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return "", errors.New("incorrect credentials")
	}

	return generateToken(user)
}

// PUBLIC
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
