package errors

import "fmt"

type ErrorType string

const (
	NotFound     ErrorType = "NOT_FOUND"
	InvalidInput ErrorType = "INVALID_INPUT"
	Internal     ErrorType = "INTERNAL"
	Unauthorized ErrorType = "UNAUTHORIZED"
)

type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// Error constructors
func NewNotFoundError(message string, err error) *AppError {
	return &AppError{
		Type:    NotFound,
		Message: message,
		Err:     err,
	}
}

func NewInvalidInputError(message string, err error) *AppError {
	return &AppError{
		Type:    InvalidInput,
		Message: message,
		Err:     err,
	}
}

func NewInternalError(message string, err error) *AppError {
	return &AppError{
		Type:    Internal,
		Message: message,
		Err:     err,
	}
}

func NewUnauthorizedError(message string, err error) *AppError {
	return &AppError{
		Type:    Unauthorized,
		Message: message,
		Err:     err,
	}
}
