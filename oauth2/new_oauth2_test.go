package oauth2

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestNewClient_TokenSourceWithInitialToken(t *testing.T) {
	initToken := &oauth2.Token{
		AccessToken: "initial-token",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(1 * time.Hour),
	}
	cfg := Config{
		ClientID:     "cid",
		ClientSecret: "csecret",
		AuthURL:      "https://auth.example.com/auth",
		TokenURL:     "https://auth.example.com/token",
		RedirectURL:  "https://app.example.com/callback",
		Scopes:       []string{"scope1", "scope2"},
	}
	client := NewClient(cfg, initToken)

	tok, err := client.Token()
	require.NoError(t, err)
	require.Equal(t, initToken, tok)
}

func TestAuthCodeURL_GeneratesCorrectURL(t *testing.T) {
	cfg := Config{
		ClientID:     "cid",
		ClientSecret: "csecret",
		AuthURL:      "https://auth.example.com/auth",
		TokenURL:     "https://auth.example.com/token",
		RedirectURL:  "https://app.example.com/callback",
		Scopes:       []string{"s1", "s2"},
	}
	client := NewClient(cfg, nil)

	urlStr := client.AuthCodeURL("mystate", oauth2.AccessTypeOffline)
	u, err := url.Parse(urlStr)
	require.NoError(t, err)
	require.Equal(t, "https", u.Scheme)
	require.Equal(t, "auth.example.com", u.Host)
	require.Equal(t, "/auth", u.Path)

	q := u.Query()
	require.Equal(t, "cid", q.Get("client_id"))
	require.Equal(t, "mystate", q.Get("state"))
	require.Equal(t, "offline", q.Get("access_type"))
}

func TestExchange_InvalidCodeReturnsError(t *testing.T) {
	cfg := Config{
		ClientID:     "cid",
		ClientSecret: "csecret",
		AuthURL:      "https://auth.example.com/auth",
		TokenURL:     "https://invalid/token",
		RedirectURL:  "https://app.example.com/callback",
	}
	client := NewClient(cfg, nil)

	_, err := client.Exchange(context.Background(), "badcode")
	require.Error(t, err)
}

// dummyRoundTripper checks that Authorization header is injected, then returns a dummy response.
type dummyRoundTripper struct {
	t *testing.T
}

func (d dummyRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	d.t.Helper()
	reqAuth := req.Header.Get("Authorization")
	require.Equal(d.t, "Bearer tok123", reqAuth)
	// Return a basic OK response
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("OK")),
		Header:     make(http.Header),
	}, nil
}

func TestHTTPClient_IncludesAuthTransport(t *testing.T) {
	initToken := &oauth2.Token{
		AccessToken: "tok123",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(1 * time.Hour),
	}
	cfg := Config{
		ClientID:     "cid",
		ClientSecret: "csecret",
		AuthURL:      "https://auth.example.com/auth",
		TokenURL:     "https://auth.example.com/token",
		RedirectURL:  "https://app.example.com/callback",
	}
	client := NewClient(cfg, initToken)
	httpClient := client.HTTPClient(context.Background())
	require.NotNil(t, httpClient)

	// Extract the oauth2.Transport to inject our dummy base RoundTripper
	trans, ok := httpClient.Transport.(*oauth2.Transport)
	require.True(t, ok, "expected Transport to be *oauth2.Transport")

	// Replace the base RoundTripper with our dummy to capture the header
	trans.Base = dummyRoundTripper{t: t}

	// Perform a request
	req, err := http.NewRequest("GET", "http://example.com/data", nil)
	require.NoError(t, err)
	resp, err := httpClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
}
