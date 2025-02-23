package handler

type ErrorResponse struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

var (
	ErrInvalidInput = &ErrorResponse{
		Code:    "INVALID_INPUT",
		Message: "The provided input is invalid",
	}
	ErrNotFound = &ErrorResponse{
		Code:    "NOT_FOUND",
		Message: "Resource not found",
	}
	ErrInternalServer = &ErrorResponse{
		Code:    "INTERNAL_ERROR",
		Message: "An internal server error occurred",
	}
)

func NewErrorResponse(code, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}

func (e *ErrorResponse) WithDetails(details map[string]interface{}) *ErrorResponse {
	e.Details = details
	return e
}
