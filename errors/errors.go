package errors

import (
	"errors"
	"go_code_snippets/errors/secure"
)

var NotFoundErr = errors.New("not found")
var ValidationErr = errors.New("validation error")

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return "invalid " + e.Field
}

func NewValidationError(public string, internal error, safeMeta map[string]any) *secure.SafeError {
	return &secure.SafeError{
		Code:     "API_VALIDATION_ERROR",
		UserMsg:  public,
		Internal: internal,
		Metadata: safeMeta,
	}
}
