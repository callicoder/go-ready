package errors

import (
	"strings"
)

// NotFoundError codes
const (
	NotFoundErrorCode = "go-ready:resource:notfound"
)

// NotFoundError holds resource not found errors
type NotFoundError struct {
	BaseError
	Cause error
}

// NewNotFoundError wraps original error with optional messages and gives NotFoundError
func NewNotFoundError(cause error, msg ...string) NotFoundError {
	nferr := NotFoundError{
		BaseError: BaseError{
			Code:        NotFoundErrorCode,
			Message:     "Resource not found",
			Description: cause.Error(),
		},
		Cause: cause,
	}
	if len(msg) > 0 {
		nferr.Message = strings.TrimSpace(strings.Join(msg, ", "))
	}
	return nferr
}
