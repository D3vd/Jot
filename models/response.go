package models

type AllNotesResponse struct {
	Count int     `json:"count"`
	Notes []Notes `json:"notes"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SimpleResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
