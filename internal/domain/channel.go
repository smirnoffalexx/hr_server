package domain

import "time"

// Channel represents a Telegram channel with channel code
type Channel struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Code      string    `json:"code" db:"code"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CreateChannelRequest represents the request to create a channel
type CreateChannelRequest struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

// GenerateChannelRequest represents the request to generate a channel code
type GenerateChannelRequest struct {
	ChannelName string `json:"channel_name" binding:"required"`
}

// GenerateBulkChannelRequest represents the request to generate multiple channel codes
type GenerateBulkChannelRequest struct {
	ChannelName string `json:"channel_name" binding:"required"`
	Count       int    `json:"count" binding:"required,min=1,max=100"`
}

// ChannelStats represents statistics for a channel
type ChannelStats struct {
	ChannelID   int    `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Code        string `json:"code"`
	UserCount   int    `json:"user_count"`
}
