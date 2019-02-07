package httpservice

import (
	"net/http"

	"github.com/callicoder/go-ready/internal/config"
)

type GoReadyTransport struct {
	Transport http.RoundTripper
	Config    config.HttpConfig
}

func (t *GoReadyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.Config.UserAgent)

	return t.Transport.RoundTrip(req)
}
