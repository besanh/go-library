package client

import (
	"strings"
	"time"
)

// A HttpClient describes an HTTP endpoint's client. This client is
// mainly used to send HTTP request based on the request configuration set into
// the client and the response handling logic configured within the client.
type HttpClient struct {
	// HTTP client to be used to execute HTTP call.
	pool *SharedPool

	// URL for this Client. This value can contain path
	// variable placeholders for substitution when sending the HTTP request
	url string

	// HTTP method
	method string

	// Name of the service to call. Together with serviceName will form the label in logger fields & error code prefixes.
	serviceName string

	// User agent value (i.e. RFC7231)
	userAgent string

	// Content MIME type
	// Default: application/json
	contentType string

	// Default request header configuration
	headers map[string]string

	timeoutAndRetryOption timeoutAndRetryOption

	// Disable request body logging
	// Default: false,
	disableReqBodyLogging bool

	// Disable response body logging
	// Default: false,
	disableRespBodyLogging bool

	// Disable request/response log redaction
	// Default: false,
	disableLogRedaction bool
}

func newHttpClient(cfg HttpClientConfig, pool *SharedPool, opts ...HttpClientOption) (*HttpClient, error) {
	client := &HttpClient{
		pool: pool,
		timeoutAndRetryOption: timeoutAndRetryOption{
			maxRetries:         defaultMaxRetries,
			maxWaitPerTry:      defaultMaxWaitPerTry * time.Second,
			maxWaitInclRetries: defaultMaxWaitInclRetries * time.Second,
			retryOnTimeout:     defaultRetryOnTimeout,
			retryOnStatusCodes: make(map[int]bool),
		},
		contentType: defaultContentType,
	}

	client.url = strings.TrimSpace(cfg.URL)
	if client.url == "" {
		return nil, ErrMissingURL
	}
	client.method = strings.TrimSpace(cfg.Method)
	if client.method == "" {
		return nil, ErrMissingMethod
	}
	client.headers = cfg.Headers

	for _, opt := range opts {
		opt(client)
	}

	if err := client.timeoutAndRetryOption.validate(); err != nil {
		return nil, err
	}
	return client, nil
}

// NewHttpClient returns a new Client instance based on the arguments
func NewHttpClient(cfg HttpClientConfig, pool *SharedPool, opts ...HttpClientOption) (*HttpClient, error) {
	return newHttpClient(cfg, pool, opts...)
}
