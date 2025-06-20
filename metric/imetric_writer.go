package metric

import (
	"fmt"
	"net/http"
)

// statusWriter captures HTTP status code.
type statusWriter struct {
	http.ResponseWriter
	statusCode int
}

type IStatusWriter interface {
	WriteHeader(code int)
	Status() string
}

// WriteHeader captures status code and delegates.
func (sw *statusWriter) WriteHeader(code int) {
	sw.statusCode = code
	sw.ResponseWriter.WriteHeader(code)
}

// Status returns captured code or default 200.
func (sw *statusWriter) Status() string {
	if sw.statusCode == 0 {
		return "200"
	}
	return fmt.Sprint(sw.statusCode)
}
