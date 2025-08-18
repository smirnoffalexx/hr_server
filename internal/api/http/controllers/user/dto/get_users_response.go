package dto

import (
	"hr-server/internal/domain"
)

type GetUsersResponse struct {
	Users []*domain.UserWithChannel `json:"users"`
}

func NewGetUsersResponse(users []*domain.UserWithChannel) *GetUsersResponse {
	return &GetUsersResponse{
		Users: users,
	}
}
