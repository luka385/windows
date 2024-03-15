package ctypes

const (
	ErrItemNotFound     = "Item not found"
	ErrInvalidParameter = "Invalid parameter"
)

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomType(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}
