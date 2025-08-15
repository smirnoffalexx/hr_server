package dto

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type GenerateChannelRequest struct {
	ChannelName string `json:"channel_name"`
}

func NewGenerateChannelRequest() *GenerateChannelRequest {
	return &GenerateChannelRequest{}
}

func (r *GenerateChannelRequest) Parse(c *gin.Context) error {
	return c.ShouldBindJSON(&r)
}

func (r *GenerateChannelRequest) Validate() error {
	err := validation.ValidateStruct(r,
		validation.Field(&r.ChannelName, validation.Required.Error("is required")),
	)
	if err != nil {
		return err
	}

	return nil
}
