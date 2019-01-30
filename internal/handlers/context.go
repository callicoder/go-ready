package handlers

import (
	"net"

	"github.com/callicoder/go-ready/internal/model"
)

type Context struct {
	RequestID string
	IPAddress net.IP
	Path      string
	UserAgent string
	Session   *model.Session
}
