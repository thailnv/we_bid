package handlers

import "net/http"

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code  int `json:"code"`
	Error `json:"error"`
}

func NewInternalServerError(err error) ErrorResponse {
	newErr := Error{
		Code:    "internal_server_error",
		Message: err.Error(),
	}

	return ErrorResponse{
		Code:  http.StatusInternalServerError,
		Error: newErr,
	}
}

func NewRecordNotFoundError(err error) ErrorResponse {
	newErr := Error{
		Code:    "record_not_found",
		Message: err.Error(),
	}

	return ErrorResponse{
		Code:  http.StatusNotFound,
		Error: newErr,
	}
}
