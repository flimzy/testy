package testy

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// ResponseHandler wraps an existing http.Response, to be served as a
// standard http.Handler
type ResponseHandler struct {
	*http.Response
}

var _ http.Handler = &ResponseHandler{}

func (h *ResponseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for header, values := range h.Header {
		for _, value := range values {
			w.Header().Add(header, value)
		}
	}
	if h.StatusCode != 0 {
		w.WriteHeader(h.StatusCode)
	}
	if h.Body != nil {
		defer h.Body.Close() // nolint: errcheck
		io.Copy(w, h.Body)
	}
}

// ServeResponse starts a test HTTP server that serves r.
func ServeResponse(r *http.Response) *httptest.Server {
	return httptest.NewServer(&ResponseHandler{r})
}
