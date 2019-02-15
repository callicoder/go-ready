package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	reqcontext "github.com/callicoder/go-ready/internal/context"
	"github.com/callicoder/go-ready/internal/model"
	"github.com/callicoder/go-ready/pkg/errors"
	"github.com/callicoder/go-ready/pkg/logger"
	"github.com/gorilla/mux"
)

const (
	headerContentType = "Content-Type"
	contentTypeJSON   = "application/json"
)

type Context interface {
	Request() *http.Request
	ResponseWriter() http.ResponseWriter
	Path() string
	Param(name string) string
	QueryParam(name string) string
	BindJSON(i interface{}) error
	JSON(statusCode int, body interface{})
	Error(err error)
	Session() *model.Session
	RequestID() string
	IPAddress() net.IP
}

type context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	pathVars       map[string]string
	requestID      string
	ipAddress      net.IP
	session        *model.Session
}

type httpError struct {
	errors.BaseError
	StatusCode int
}

func NewContext(w http.ResponseWriter, r *http.Request) Context {
	return &context{
		request:        r,
		responseWriter: w,
		requestID:      reqcontext.RequestID(r.Context()),
		ipAddress:      reqcontext.IPAddress(r.Context()),
		session:        reqcontext.Session(r.Context()),
		pathVars:       mux.Vars(r),
	}
}

func (c *context) Request() *http.Request {
	return c.request
}

func (c *context) ResponseWriter() http.ResponseWriter {
	return c.responseWriter
}

func (c *context) Path() string {
	return c.request.URL.Path
}

func (c *context) Param(name string) string {
	return c.pathVars[name]
}

func (c *context) QueryParam(name string) string {
	query := c.request.URL.Query()
	return query.Get(name)
}

func (c *context) Session() *model.Session {
	return c.session
}

func (c *context) RequestID() string {
	return c.requestID
}

func (c *context) IPAddress() net.IP {
	return c.ipAddress
}

func (c *context) Error(err error) {
	var httpErr httpError

	switch err.(type) {
	case errors.InternalError:
		inErr := err.(errors.InternalError)
		httpErr = httpError{
			BaseError:  inErr.BaseError,
			StatusCode: http.StatusInternalServerError,
		}

	case errors.NotFoundError:
		nfErr := err.(errors.NotFoundError)
		httpErr = httpError{
			BaseError:  nfErr.BaseError,
			StatusCode: http.StatusNotFound,
		}

	case errors.ValidationError:
		vErr := err.(errors.ValidationError)
		httpErr = httpError{
			BaseError:  vErr.BaseError,
			StatusCode: http.StatusBadRequest,
		}

	case errors.UnauthorizedError:
		aErr := err.(errors.UnauthorizedError)
		httpErr = httpError{
			BaseError:  aErr.BaseError,
			StatusCode: http.StatusUnauthorized,
		}

	default:
		unexpectedErr := errors.NewInternalError(err)
		httpErr = httpError{
			BaseError:  unexpectedErr.BaseError,
			StatusCode: http.StatusInternalServerError,
		}
	}

	c.JSON(httpErr.StatusCode, httpErr.BaseError)
}

func (c *context) JSON(statusCode int, body interface{}) {
	c.responseWriter.Header().Set(headerContentType, contentTypeJSON)
	if body == nil {
		c.responseWriter.WriteHeader(statusCode)
		return
	}

	if err := json.NewEncoder(c.responseWriter).Encode(body); err != nil {
		logger.Errorf("Failed to write response: %v", err)
	}
}

func (c *context) BindJSON(i interface{}) error {
	req := c.request
	if err := json.NewDecoder(req.Body).Decode(i); err != nil {
		if ute, ok := err.(*json.UnmarshalTypeError); ok {
			return errors.Wrap(err, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset))
		} else if se, ok := err.(*json.SyntaxError); ok {
			return errors.Wrap(err, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error()))
		}
		return err
	}
	return nil
}
