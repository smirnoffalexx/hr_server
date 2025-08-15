package dto

import (
	"fmt"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type GenerateBulkChannelRequest struct {
	ChannelNames []string `json:"channel_names"`
}

func NewGenerateBulkChannelRequest() *GenerateBulkChannelRequest {
	return &GenerateBulkChannelRequest{}
}

func (r *GenerateBulkChannelRequest) Parse(c *gin.Context) error {
	return c.ShouldBindJSON(&r)
}

func (r *GenerateBulkChannelRequest) Validate() error {
	err := validation.ValidateStruct(r,
		validation.Field(&r.ChannelNames,
			validation.Required.Error("channel names array is required"),
			validation.Length(1, 100).Error("must have between 1 and 100 channel names"),
		),
	)
	if err != nil {
		return err
	}

	// Validate each channel name individually
	for i, name := range r.ChannelNames {
		if err := validation.Validate(name, validation.Required.Error("channel name cannot be empty")); err != nil {
			return fmt.Errorf("channel name at index %d: %w", i, err)
		}
	}

	return nil
}
