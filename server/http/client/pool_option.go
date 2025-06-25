package client

import (
	"log"
	"net/http"
	"net/url"
)

type PoolOption func(c *http.Client, t *http.Transport)

func WithProxy(proxyURL string) PoolOption {
	return func(c *http.Client, t *http.Transport) {
		if proxyURL == "" {
			t.Proxy = nil
			return
		}
		parsedURL, err := url.Parse(proxyURL)
		if err != nil {
			log.Fatalf("failed to parse proxy URL. Err: %v", err)
			return
		}

		t.Proxy = http.ProxyURL(parsedURL)
	}
}
