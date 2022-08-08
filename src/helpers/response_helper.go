package helpers

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func BuildResponse(status bool, message string, err string, data interface{}) Response {
	res := Response{
		Status:  true,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}
