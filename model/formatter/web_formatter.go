package formatter

import "github.com/sandriansyafridev/golang/api/book/model/response"

func BuildResponseSuccess(code int, data interface{}) response.WebResponse {
	return response.WebResponse{
		Code:    code,
		Status:  true,
		Message: "ok",
		Errors:  nil,
		Data:    data,
	}
}

func BuildResponseError(code int, errors interface{}) response.WebResponse {
	return response.WebResponse{
		Code:    code,
		Status:  false,
		Message: "!ok",
		Errors:  errors,
		Data:    nil,
	}
}
