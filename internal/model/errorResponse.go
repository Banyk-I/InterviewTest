package model

type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"status_code"`
}
