package response

import (
	"net/http"
	"strings"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

func SuccessReponse(msg string, data ...interface{}) *Response {
	return &Response{
		StatusCode: http.StatusOK,
		Message:    msg,
		Data:       data,
		Error:      nil,
	}
}
func ErrorResponse(code int, msg string, err string, data ...interface{}) *Response {
	errSplit := strings.Split(err, "\n")
	return &Response{
		StatusCode: code,
		Message:    msg,
		Data:       data,
		Error:      errSplit,
	}
}
