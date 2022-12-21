package helper

import "strings"

type CommonResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"errors"`
}

type EmptyObject struct {
}

func CreateCommonResponse(status bool, message string, data interface{}) CommonResponse {
	res := CommonResponse{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func CreateCommonErrorResponse(message string, error string, data interface{}) CommonResponse {
	splitError := strings.Split(error, "\n")
	res := CommonResponse{
		Status:  false,
		Message: message,
		Errors:  splitError,
		Data:    data,
	}
	return res

}
