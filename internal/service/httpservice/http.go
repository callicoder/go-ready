package httpservice

import (
	"net/http"

	"github.com/callicoder/go-ready/internal/config"
)

type HttpService interface {
	MakeClient() *http.Client
}

type httpService struct {
	config config.HttpConfig
}

func NewHttpService(config config.HttpConfig) HttpService {
	return &httpService{config}
}

func (h *httpService) MakeClient() *http.Client {
	return NewHttpClient(h.config)
}
