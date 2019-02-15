package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/callicoder/go-ready/internal/context"
	"github.com/callicoder/go-ready/internal/model"
	"github.com/callicoder/go-ready/pkg/errors"
	"github.com/callicoder/go-ready/pkg/logger"
)

const (
	headerContentType = "Content-Type"
	contentTypeJSON   = "application/json"
)

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	RequestID      string
	IPAddress      net.IP
	Session        *model.Session
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request:        r,
		ResponseWriter: w,
		RequestID:      context.RequestID(r.Context()),
		IPAddress:      context.IPAddress(r.Context()),
		Session:        context.Session(r.Context()),
	}
}

type httpError struct {
	errors.BaseError
	StatusCode int
}

func (c *Context) Error(err error) {
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

func (c *Context) JSON(statusCode int, body interface{}) {
	c.ResponseWriter.Header().Set(headerContentType, contentTypeJSON)
	if body == nil {
		c.ResponseWriter.WriteHeader(statusCode)
		return
	}

	if err := json.NewEncoder(c.ResponseWriter).Encode(body); err != nil {
		logger.Errorf("Failed to write response: %v", err)
	}
}

func (c *Context) BindJSON(i interface{}) error {
	req := c.Request
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
