# Fingerprint Package 

The `fingerprint` package provides utilities to safely extract, sanitize, and format HTTP request headers for logging and tracing purposes. It helps prevent log injection attacks and ensures standardized request tracking across services.

## Architecture Flow

1. **Incoming Request**: An `http.Request` is received by an HTTP handler or middleware.
2. **Extraction & Sanitization**: The package extracts standard identifying headers (User-Agent, Accept-Language, Referer, etc.) while sanitizing outputs to replace invalid/control characters with safe alternatives (like `` or `􏿮`).
3. **Formatting**: Provides methods to format the extracted request data into either a structured single-line text format or a Markdown-formatted list for easy readability.

## Usage

```go
import "github.com/besanh/go-library/fingerprint"
```

### 1. Generating a Single-Line Fingerprint
Use `RequestFingerprint` to log a comprehensive, sanitized summary of the HTTP request headers.

```go
func middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        logLine := fingerprint.RequestFingerprint(r)
        // logLine: " en-US Mozilla/5.0... R=https://... A=*/* E=gzip Connection=keep-alive"
        log.Println(logLine)
        next.ServeHTTP(w, r)
    })
}
```

### 2. Generating a Markdown Fingerprint
Use `FingerprintMD` to generate a Markdown bulleted list, useful for including request context in error reports or issue trackers.

```go
mdReport := fingerprint.FingerprintMD(r)
/* Outputs:
- **IP**: 192.168.1.1
- **Accept-Language**: en-US
- **User-Agent**: ...
*/
```

### 3. Safely Extracting Specific Headers
Use `SafeHeader` to safely extract and sanitize a single header value.

```go
userAgent := fingerprint.SafeHeader(r, "User-Agent")
```
