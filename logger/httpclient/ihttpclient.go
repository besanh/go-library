package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

func (h *httpHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	payload := map[string]any{
		"time":  time.Now().Format(time.RFC3339Nano),
		"level": level.String(),
		"msg":   msg,
	}
	buf := &bytes.Buffer{}
	json.NewEncoder(buf).Encode(payload)

	req, _ := http.NewRequest("POST", h.url, buf)
	req.Header.Set("Content-Type", "application/json")
	h.client.Do(req)
}
