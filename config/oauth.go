package config

import (
	"os"
)

type OAuth struct {
	BaseUrl      string
	ClientID     string
	ClientSecret string
	RedirectUrl  string
	ScopeOpenId  string
	ScopeProfile string
	ScopeEmail   string
}

func newOAuth() *OAuth {
	return &OAuth{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		RedirectUrl:  os.Getenv("OAUTH_GOOGLE_REDIRECT_URL"),
	}
}

func (c OAuth) GenerateOAuthTokenRequestURL() string {
	baseUrl := "https://api.instagram.com/oauth/access_token"
	return baseUrl
}
