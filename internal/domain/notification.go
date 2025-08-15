package domain

import "time"

// Notification represents a notification message
type Notification struct {
	ID        int        `json:"id" db:"id"`
	Message   string     `json:"message" db:"message"`
	ImageURL  *string    `json:"image_url,omitempty" db:"image_url"`
	Emoji     *string    `json:"emoji,omitempty" db:"emoji"`
	Status    string     `json:"status" db:"status"`
	SentAt    *time.Time `json:"sent_at,omitempty" db:"sent_at"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}

// SendNotificationRequest represents the request to send a notification
type SendNotificationRequest struct {
	Message  string  `json:"message" binding:"required"`
	ImageURL *string `json:"image_url,omitempty"`
	Emoji    *string `json:"emoji,omitempty"`
}

// TelegramMessage represents a Telegram bot message
type TelegramMessage struct {
	ChatID      int64  `json:"chat_id"`
	Text        string `json:"text"`
	ParseMode   string `json:"parse_mode,omitempty"`
	ReplyMarkup string `json:"reply_markup,omitempty"`
}
