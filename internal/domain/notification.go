package domain

// NotificationData represents the data to send a notification
type NotificationData struct {
	Message  string  `json:"message"`
	ImageURL *string `json:"image_url,omitempty"`
}
