package presenter

type ApiError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewApiError(statusCode int, message string) *ApiError {
	return &ApiError{
		StatusCode: statusCode,
		Message:    message,
	}
}
