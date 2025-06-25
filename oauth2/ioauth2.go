package oauth2

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type IOAuth2 interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
	Token() (*oauth2.Token, error)
	HTTPClient(ctx context.Context) *http.Client
}

// AuthCodeURL returns the URL to redirect users for OAuth2 authorization.
func (c *Client) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return c.config.AuthCodeURL(state, opts...)
}

// Exchange exchanges an authorization code for a token.
func (c *Client) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return c.config.Exchange(ctx, code)
}

// Token retrieves a valid token, refreshing if necessary.
func (c *Client) Token() (*oauth2.Token, error) {
	return c.tokenSource.Token()
}

// HTTPClient returns an *http.Client which automatically injects OAuth2 tokens in requests.
func (c *Client) HTTPClient(ctx context.Context) *http.Client {
	return oauth2.NewClient(ctx, c.tokenSource)
}
