package oauth2

import (
	"golang.org/x/oauth2"
)

type (
	OAuth2 struct {
		Config OAuth2Config
	}

	OAuth2Config struct {
		ClientId     string
		ClientSecret string
		Scopes       []string
		Endpoint     oauth2.Endpoint
		Redirect     string
	}
)

func NewOAuth2(config OAuth2Config) IOAuth2 {
	return &OAuth2{
		Config: config,
	}
}
