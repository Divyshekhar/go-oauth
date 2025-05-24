package utils

import (
	"context"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

var OauthConfig *oauth2.Config

func InitOAuth() error {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return err
	}
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		return err
	}
	OauthConfig = config
	return nil
}

func GetURL() string {
	return OauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.SetAuthURLParam("prompt", "consent"))
}

func Exchange(code string) *oauth2.Token {
	token, _ := OauthConfig.Exchange(context.Background(), code)
	return token
}
