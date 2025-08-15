package channel

import (
	"fmt"
	"hr-server/internal/api/http/controllers/channel/dto"
	"hr-server/internal/api/http/controllers/common"
	"hr-server/internal/register"
	"hr-server/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ChannelController struct {
	channelService *service.ChannelService
}

func NewChannelController(sr *register.StorageRegister) *ChannelController {
	return &ChannelController{sr.ChannelService()}
}

// GenerateChannel godoc
// @Summary Generate a new channel code
// @Description Generate a new channel code for a specific channel
// @Tags Channels
// @Accept json
// @Produce json
// @Param request body dto.GenerateChannelRequest true "Generate channel request"
// @Success 200 {object} domain.Channel
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security XAuthToken
// @Router /channel/generate [post]
func (c *ChannelController) GenerateChannelHandler(sr *register.StorageRegister) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := dto.NewGenerateChannelRequest()
		if err := req.Parse(ctx); err != nil {
			logrus.Error("unable to parse a request: ", err)
			ctx.JSON(http.StatusBadRequest, common.ErrorResponse{Error: err.Error()})
			return
		}

		if err := req.Validate(); err != nil {
			logrus.Error("error of validation: ", err)
			ctx.JSON(http.StatusBadRequest, common.ErrorResponse{Error: err.Error()})
			return
		}

		channel, err := c.channelService.GenerateChannel(req.ChannelName)
		if err != nil {
			logrus.Error("error while generate channel: ", err)
			ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: fmt.Sprintf("failed to generate channel: %v", err)})
			return
		}

		ctx.JSON(http.StatusOK, channel)
	}
}

// GenerateBulkChannel godoc
// @Summary Generate multiple channel codes
// @Description Generate multiple channel codes for different channel names
// @Tags Channels
// @Accept json
// @Produce json
// @Param request body dto.GenerateBulkChannelRequest true "Generate bulk channel request with array of names"
// @Success 200 {array} domain.Channel
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security XAuthToken
// @Router /channel/bulk [post]
func (c *ChannelController) GenerateBulkChannelHandler(sr *register.StorageRegister) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := dto.NewGenerateBulkChannelRequest()
		if err := req.Parse(ctx); err != nil {
			logrus.Error("unable to parse a request: ", err)
			ctx.JSON(http.StatusBadRequest, common.ErrorResponse{Error: err.Error()})
			return
		}

		if err := req.Validate(); err != nil {
			logrus.Error("error of validation: ", err)
			ctx.JSON(http.StatusBadRequest, common.ErrorResponse{Error: err.Error()})
			return
		}

		channels, err := c.channelService.GenerateBulkChannel(req.ChannelNames)
		if err != nil {
			logrus.Error("error while generate bulk channel: ", err)
			ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: fmt.Sprintf("failed to generate bulk channel: %v", err)})
			return
		}

		ctx.JSON(http.StatusOK, channels)
	}
}

// GetChannelByCode godoc
// @Summary Get channel by code
// @Description Get channel information by its code
// @Tags Channels
// @Accept json
// @Produce json
// @Param code path string true "Channel code"
// @Success 200 {object} domain.Channel
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security XAuthToken
// @Router /channel/{code} [get]
func (c *ChannelController) GetChannelByCodeHandler(sr *register.StorageRegister) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.Param("code")
		if code == "" {
			ctx.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Channel code is required"})
			return
		}

		channel, err := c.channelService.GetChannelByCode(code)
		if err != nil {
			logrus.Error("error while get channel by code: ", err)
			ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: fmt.Sprintf("failed to get channel by code '%s': %v", code, err)})
			return
		}

		if channel == nil {
			ctx.JSON(http.StatusNotFound, common.ErrorResponse{Error: "Channel code not found"})
			return
		}

		ctx.JSON(http.StatusOK, channel)
	}
}

// GetChannels godoc
// @Summary Get all channels
// @Description Get all Telegram channels
// @Tags Channels
// @Accept json
// @Produce json
// @Success 200 {array} domain.Channel
// @Failure 500 {object} common.ErrorResponse
// @Security XAuthToken
// @Router /channels [get]
func (c *ChannelController) GetChannelsHandler(sr *register.StorageRegister) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		channels, err := c.channelService.GetAll()
		if err != nil {
			logrus.Error("error while get all channels: ", err)
			ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: fmt.Sprintf("failed to get all channels: %v", err)})
			return
		}

		response := dto.NewGetChannelsResponse(channels)
		ctx.JSON(http.StatusOK, response)
	}
}
