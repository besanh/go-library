# OAuth2 Package

The `oauth2` package provides a standardized encapsulation of `golang.org/x/oauth2`. It manages complex OAuth2 client authorization code flows, token exchange sequences, and authenticated HTTP clients effortlessly.

## Architecture Flow

1. **Configuration & Endpoint Setup**: Consolidates credentials (Client ID/Secret) and endpoints (Auth URL, Token URL) into a simplified `Config` model.
2. **Token Management**: Wrapping the standard `oauth2.TokenSource`, the client can execute exchanges from callback codes to retrieve access tokens, and natively supports background refreshing of expired access tokens using underlying refresh tokens.
3. **Transport Injection**: Provides a pre-configured HTTP client that automatically seamlessly injects the managed Bearer tokens into outbound HTTP requests.

## Usage

```go
import "github.com/besanh/go-library/oauth2"
```

### 1. Initialization

Set up an OAuth2 Client instance with the required Identity Provider endpoints.

```go
cfg := oauth2.Config{
    ClientID:     "my-app-id",
    ClientSecret: "my-app-secret",
    AuthURL:      "https://idp.example.com/auth",
    TokenURL:     "https://idp.example.com/token",
    RedirectURL:  "https://my.app.com/callback",
    Scopes:       []string{"openid", "profile", "email"},
}

// Create a new client (initialToken can be nil if starting a fresh flow)
client := oauth2.NewClient(cfg, nil)
```

### 2. The Authorization Code Flow

Redirect the user to the IDP, catch the callback, and exchange it for a token.

```go
// 1. Get the redirect URL and send the user there
state := "random-state-string"
url := client.AuthCodeURL(state)
http.Redirect(w, r, url, http.StatusTemporaryRedirect)

// ... Later inside your callback handler ...

// 2. Extract the code from the query parameters
code := r.URL.Query().Get("code")

// 3. Exchange the code for a Token 
token, err := client.Exchange(r.Context(), code)
if err != nil {
    // Handle exchange failure
}
```

### 3. Authenticated HTTP Calls

Easily perform HTTP requests strictly using the acquired and refreshed token.

```go
// Creates an *http.Client utilizing the internal TokenSource
httpClient := client.HTTPClient(context.Background())

// Authorization: Bearer <token> is injected automatically
resp, err := httpClient.Get("https://api.example.com/protected-resource")
```
