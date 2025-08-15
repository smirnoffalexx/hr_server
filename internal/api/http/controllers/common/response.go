package common

type SuccessResponse struct{}

type ErrorResponse struct {
	Error string `json:"error"`
}
