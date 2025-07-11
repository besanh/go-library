package httpclient

import (
	"bytes"
	"errors"
	"net/http"
	"time"
)

type HttpHookWriter struct {
	URL string
}

func (h *HttpHookWriter) Write(p []byte) (n int, err error) {
	req, err := http.NewRequest("POST", h.URL, bytes.NewBuffer(p))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return 0, errors.New("http hook responded with status: " + resp.Status)
	}
	return len(p), nil
}
