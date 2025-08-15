package dto

import (
	"hr-server/internal/domain"
)

type GetChannelsResponse struct {
	Channels []*domain.Channel `json:"channels"`
}

func NewGetChannelsResponse(channels []*domain.Channel) *GetChannelsResponse {
	return &GetChannelsResponse{
		Channels: channels,
	}
}
