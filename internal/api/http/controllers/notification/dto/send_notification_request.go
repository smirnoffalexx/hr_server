package dto

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type SendNotificationRequest struct {
	Message  string  `json:"message"`
	ImageURL *string `json:"image_url,omitempty"`
}

func NewSendNotificationRequest() *SendNotificationRequest {
	return &SendNotificationRequest{}
}

func (r *SendNotificationRequest) Parse(c *gin.Context) error {
	return c.ShouldBindJSON(&r)
}

func (r *SendNotificationRequest) Validate() error {
	err := validation.ValidateStruct(r,
		validation.Field(&r.Message, validation.Required.Error("is required")),
		validation.Field(&r.ImageURL, validation.By(validateImageURL)),
	)
	if err != nil {
		return err
	}

	return nil
}

// validateImageURL validates that the image URL is properly formatted
func validateImageURL(value interface{}) error {
	if value == nil {
		return nil // Optional field
	}

	urlStr, ok := value.(*string)
	if !ok {
		return fmt.Errorf("image_url must be a string pointer")
	}

	if urlStr == nil {
		return nil // Nil pointer is valid (optional)
	}

	// Trim whitespace
	trimmed := strings.TrimSpace(*urlStr)
	if trimmed == "" {
		return fmt.Errorf("image_url cannot be empty if provided")
	}

	// Parse URL to ensure it's valid
	parsedURL, err := url.Parse(trimmed)
	if err != nil {
		return fmt.Errorf("image_url has invalid URL format: %v", err)
	}

	// Check if it's HTTP or HTTPS
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("image_url must be HTTP or HTTPS URL")
	}

	// Check if host is provided
	if parsedURL.Host == "" {
		return fmt.Errorf("image_url must have a valid host")
	}

	return nil
}
