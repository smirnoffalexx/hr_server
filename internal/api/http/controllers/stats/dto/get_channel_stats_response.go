package dto

import (
	"hr-server/internal/domain"
)

type GetChannelStatsResponse struct {
	Stats *domain.ChannelStats `json:"stats"`
}

func NewGetChannelStatsResponse(stats *domain.ChannelStats) *GetChannelStatsResponse {
	return &GetChannelStatsResponse{
		Stats: stats,
	}
}
