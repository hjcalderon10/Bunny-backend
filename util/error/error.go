package error

import (
	"fmt"
)

var InternalServerError = fmt.Errorf("Unexpected error")

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (errMsg ErrorMessage) Error() string {
	return fmt.Sprintf("%s: %s", errMsg.Code, errMsg.Message)
}

func HandleError(code int, message string) *ErrorMessage {
	return &ErrorMessage{
		code,
		message,
	}
}
