package dto

import (
	"hr-server/internal/domain"
)

type GetUsersResponse struct {
	Users []*domain.User `json:"users"`
}

func NewGetUsersResponse(users []*domain.User) *GetUsersResponse {
	return &GetUsersResponse{
		Users: users,
	}
}
