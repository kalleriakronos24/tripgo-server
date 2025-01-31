package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/kalleriakronos24/khaimal-group/config"
)

type GoogleOauthToken struct {
	Access_token string
	Id_token     string
}

type GoogleUserResult struct {
	Id             string
	Email          string
	Verified_email bool
	Name           string
	Given_name     string
	Family_name    string
	Picture        string
	Locale         string
}

func GetGoogleOauthToken(code string) (*GoogleOauthToken, error) {
	const rootURl = "https://oauth2.googleapis.com/token"

	config := config.AppConfig
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", config.GOOGLE_OAUTH_CLIENT_ID)
	values.Add("client_secret", config.GOOGLE_OAUTH_CLIENT_SECRET)
	values.Add("redirect_uri", config.GOOGLE_OAUTH_REDIRECT_URL)

	query := values.Encode()

	req, err := http.NewRequest("POST", rootURl, bytes.NewBufferString(query))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve token")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleOauthTokenRes map[string]interface{}

	if err := json.Unmarshal(resBody, &GoogleOauthTokenRes); err != nil {
		return nil, err
	}

	tokenBody := &GoogleOauthToken{
		Access_token: GoogleOauthTokenRes["access_token"].(string),
		Id_token:     GoogleOauthTokenRes["id_token"].(string),
	}

	return tokenBody, nil
}

type GoogleUserRes struct {
	ID             string
	Email          string
	Verified_email bool
	Name           string
	Given_name     string
	Picture        string
	Locale         string
}

func GetGoogleUser(access_token string, id_token string) (*GoogleUserResult, error) {
	rootUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", access_token)

	req, err := http.NewRequest("GET", rootUrl, nil)
	if err != nil {
		return nil, err
	}

	log.Printf("ID TOKEN >>>> %v", id_token)

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", id_token))

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var googleUserResponse GoogleUserResult

	if err := json.Unmarshal(resBody, &googleUserResponse); err != nil {
		return nil, err
	}

	log.Printf("GOOGLE USER >>> %v", googleUserResponse)

	userBody := &GoogleUserResult{
		Id:             googleUserResponse.Id,
		Email:          googleUserResponse.Email,
		Verified_email: googleUserResponse.Verified_email,
		Name:           googleUserResponse.Name,
		Given_name:     googleUserResponse.Given_name,
		Picture:        googleUserResponse.Picture,
		Locale:         googleUserResponse.Locale,
	}

	return userBody, nil
}
