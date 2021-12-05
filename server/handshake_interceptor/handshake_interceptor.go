package handshake_interceptor

import (
	"net/http"
	"net/url"
)

type Interface interface {
	 Intercept(query url.Values, headers http.Header) bool
}
