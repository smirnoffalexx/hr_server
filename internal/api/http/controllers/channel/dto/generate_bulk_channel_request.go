package dto

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type GenerateBulkChannelRequest struct {
	ChannelName string `json:"channel_name"`
	Count       int    `json:"count"`
}

func NewGenerateBulkChannelRequest() *GenerateBulkChannelRequest {
	return &GenerateBulkChannelRequest{}
}

func (r *GenerateBulkChannelRequest) Parse(c *gin.Context) error {
	return c.ShouldBindJSON(&r)
}

func (r *GenerateBulkChannelRequest) Validate() error {
	err := validation.ValidateStruct(r,
		validation.Field(&r.ChannelName, validation.Required.Error("is required")),
		validation.Field(&r.Count, validation.Required.Error("is required"), validation.Min(1).Error("must be at least 1"), validation.Max(100).Error("must be at most 100")),
	)
	if err != nil {
		return err
	}

	return nil
}
