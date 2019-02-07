package httpservice

import (
	"net"
	"net/http"
	"time"

	"github.com/callicoder/go-ready/internal/config"
)

func NewTransport(config config.HttpConfig) http.RoundTripper {
	dialContext := (&net.Dialer{
		Timeout:   time.Duration(config.ConnectTimeoutSec) * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext

	return &GoReadyTransport{
		Config: config,
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           dialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   time.Duration(config.ConnectTimeoutSec) * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}

func NewHttpClient(config config.HttpConfig) *http.Client {
	return &http.Client{
		Transport: NewTransport(config),
		Timeout:   time.Duration(config.RequestTimeoutSec) * time.Second,
	}
}
