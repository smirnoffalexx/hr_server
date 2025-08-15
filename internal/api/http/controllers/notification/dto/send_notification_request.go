package dto

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type SendNotificationRequest struct {
	Message  string  `json:"message"`
	ImageURL *string `json:"image_url,omitempty"`
	Emoji    *string `json:"emoji,omitempty"`
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
	)
	if err != nil {
		return err
	}

	return nil
}
