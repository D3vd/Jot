package models

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SimpleResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
