package dto

type GetStatsResponse struct {
	Message string `json:"message"`
}

func NewGetStatsResponse(message string) *GetStatsResponse {
	return &GetStatsResponse{
		Message: message,
	}
}
