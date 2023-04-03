package models

import "net/http"

func StatusBadRequest(errors []string) Response {
	return Response{
		Code:    http.StatusBadRequest,
		Message: "bad request",
		Errors:  errors,
		Data:    nil,
	}
}

func StatusBadGateway(errors []string) Response {
	return Response{
		Code:    http.StatusBadGateway,
		Message: "bad gateway",
		Errors:  errors,
		Data:    nil,
	}
}

func StatusUnauthorized(errors []string) Response {
	return Response{
		Code:    http.StatusUnauthorized,
		Message: "unauthorized",
		Errors:  errors,
		Data:    nil,
	}
}

func StatusCreated(data interface{}) Response {
	return Response{
		Code:    http.StatusCreated,
		Message: "created",
		Errors:  []string{},
		Data:    data,
	}
}

func StatusOk(data interface{}) Response {
	return Response{
		Code:    http.StatusOK,
		Message: "ok",
		Errors:  []string{},
		Data:    data,
	}
}
