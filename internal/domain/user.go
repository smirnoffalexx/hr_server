package domain

import "time"

// User represents a Telegram user
type User struct {
	ID         int       `json:"id"`
	TelegramID int64     `json:"telegram_id"`
	Username   string    `json:"username"`
	ChannelID  *int      `json:"channel_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// UserWithChannel represents a Telegram user with channel information
type UserWithChannel struct {
	ID          int       `json:"id"`
	TelegramID  int64     `json:"telegram_id"`
	Username    string    `json:"username"`
	ChannelID   *int      `json:"channel_id"`
	ChannelName *string   `json:"channel_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
