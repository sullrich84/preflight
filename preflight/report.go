package preflight

import "net/http"

type Report struct {
	Request        *http.Request
	Response       *http.Response
	AllowedOrigins string
	AllowedMethods string
	AllowedHeaders string
}
