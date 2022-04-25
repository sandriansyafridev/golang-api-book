package formatter

import (
	"fmt"
	"strings"

	"github.com/sandriansyafridev/golang/api/book/model/response"
)

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

	str := fmt.Sprintf("%v", errors)
	strSplit := strings.Split(str, "\n")

	if len(strSplit) != 1 {
		errors = strSplit
	}

	return response.WebResponse{
		Code:    code,
		Status:  false,
		Message: "!ok",
		Errors:  errors,
		Data:    nil,
	}
}
