package dto

var (
	ErrorCodeNotFound       ErrorCode = ErrorCode{Name: "NOT_FOUND", Value: 404}
	ErrorInvalidRequest     ErrorCode = ErrorCode{Name: "INVALID_REQUEST", Value: 400}
	ErrorCodeInternalServer ErrorCode = ErrorCode{Name: "INTERNAL_SERVER_ERROR", Value: 500}
)

type ErrorCode struct {
	Name  string
	Value int
}

type ServiceError struct {
	Code    ErrorCode `json:"code,omitempty"`
	Message string    `json:"message,omitempty"`
}

func (e ServiceError) Error() string {
	return e.Message
}
