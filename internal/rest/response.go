package rest

import (
	"net/http"
)

type Response struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func InternalServerErrorResponse() *Response {
	return &Response{
		Code:  http.StatusInternalServerError,
		Error: "Internal Server Error",
	}
}
