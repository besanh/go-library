// oauth2/client.go
package oauth2

import (
	"context"

	"golang.org/x/oauth2"
)

// Config holds OAuth2 parameters for creating a client.
type Config struct {
	ClientID     string   // OAuth2 client ID
	ClientSecret string   // OAuth2 client secret
	AuthURL      string   // Authorization endpoint
	TokenURL     string   // Token endpoint
	RedirectURL  string   // Redirect URL for auth code flow
	Scopes       []string // OAuth2 scopes
}

// Client wraps an oauth2.Config and TokenSource to provide HTTP clients and token management.
type Client struct {
	config      *oauth2.Config
	tokenSource oauth2.TokenSource
}

// NewClient creates a new OAuth2 Client. If initialToken is nil, TokenSource will fetch tokens via Exchange().
func NewClient(cfg Config, initialToken *oauth2.Token) *Client {
	oauthCfg := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Scopes:       cfg.Scopes,
		RedirectURL:  cfg.RedirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  cfg.AuthURL,
			TokenURL: cfg.TokenURL,
		},
	}
	// Create TokenSource, which will auto-refresh as needed
	ts := oauthCfg.TokenSource(context.Background(), initialToken)
	return &Client{
		config:      oauthCfg,
		tokenSource: ts,
	}
}
