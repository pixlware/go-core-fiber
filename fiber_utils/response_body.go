package fiber_utils

import "net/http"

type ResponseBody struct {
	StatusCode    int    `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	Message       string `json:"message"`
	Data          any    `json:"data,omitempty"`
	Error         any    `json:"error,omitempty"`
	ErrorCode     string `json:"errorCode,omitempty"`
}

func NewSuccessResponseBody(statusCode int, message string, data any) *ResponseBody {
	return &ResponseBody{
		StatusCode:    statusCode,
		StatusMessage: http.StatusText(statusCode),
		Message:       message,
		Data:          data,
	}
}

func NewErrorResponseBody(statusCode int, message string, error any, errorCode string) *ResponseBody {
	return &ResponseBody{
		StatusCode:    statusCode,
		StatusMessage: http.StatusText(statusCode),
		Message:       message,
		Error:         error,
		ErrorCode:     errorCode,
	}
}
