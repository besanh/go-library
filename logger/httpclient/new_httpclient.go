package httpclient

import (
	"net/http"

	"github.com/rs/zerolog"
)

type IHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type httpHook struct {
	url    string
	client IHTTPClient
}

func NewHTTPHook(url string, client IHTTPClient) zerolog.Hook {
	if client == nil {
		client = http.DefaultClient
	}
	return &httpHook{url: url, client: client}
}
