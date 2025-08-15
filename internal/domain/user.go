package domain

import "time"

// User represents a Telegram user
type User struct {
	ID         int       `json:"id" db:"id"`
	TelegramID int64     `json:"telegram_id" db:"telegram_id"`
	Username   string    `json:"username" db:"username"`
	ChannelID  *int      `json:"channel_id" db:"channel_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Channel    *Channel  `json:"channel,omitempty"`
}
