package dto

type GetChannelsStatsResponse struct {
	Message string `json:"message"`
}

func NewGetChannelsStatsResponse(message string) *GetChannelsStatsResponse {
	return &GetChannelsStatsResponse{
		Message: message,
	}
}
