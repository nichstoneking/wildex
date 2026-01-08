package models

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func NewSuccessResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(message string, err string) *ErrorResponse {
	return &ErrorResponse{
		Status:  "error",
		Message: message,
		Error:   err,
	}
}
