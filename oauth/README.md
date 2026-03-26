# OAuth JWT Verification Package

The `oauth` package provides robust and secure JWT verification using public keys dynamically fetched from an external JWKS URL. It guarantees that incoming requests are authenticated by a trusted Identity Provider using Asymmetric RSA cryptographic signing.

## Architecture Flow

1. **Initialization**: The `Authenticator` is initialized with a required `jwksURL`. Upon startup, it fetches the JWKS (JSON Web Key Set) directly from the Identity Provider.
2. **Background Refresh**: The `keyfunc` library handles automatic, periodic background refreshes (configurable via options) to ensure key rotation policies won't disrupt authentication.
3. **Middleware Interception**: The provided standard middleware retrieves the `Bearer` token from the `Authorization` header on incoming requests.
4. **Verification**: The token is parsed and its signature verified securely against the cached public keys. It enforces `alg: RS256` explicitly.
5. **Context Injection**: Successfully validated tokens are unpacked, the `sub` (User ID) is extracted, and securely injected into the `context.Context` to be utilized by downstream business logic.

## Usage

```go
import "github.com/besanh/go-library/oauth"
```

### 1. Setting up the Authenticator

Initialize the authenticator when the application starts. Ensure clean shutdown of the background key refreshing loop via `Close()`.

```go
authenticator, err := oauth.NewAuthenticator(
    oauth.WithJWKSURL("https://my-idp.example.com/.well-known/jwks.json"),
    oauth.WithRefreshTimeout(5 * time.Minute),
)
if err != nil {
    log.Fatal("Failed to setup OAuth:", err)
}

// Clean up background task on exit
defer authenticator.(*oauth.Authenticator).Close()
```

### 2. Using the Middleware

Wrap your HTTP handlers or routers to secure specific endpoints automatically.

```go
mux := http.NewServeMux()
mux.HandleFunc("/api/private", myPrivateHandler)

secureMux := authenticator.(*oauth.Authenticator).StandardMiddleware(mux)

http.ListenAndServe(":8080", secureMux)
```

### 3. Extracting Context in Handlers

Once requests pass the middleware, retrieve the injected User ID downstream.

```go
import "github.com/besanh/go-library/oauth/context"

func myPrivateHandler(w http.ResponseWriter, r *http.Request) {
    userID := context.GetUserID(r.Context())
    // Process request for userID...
}
```
