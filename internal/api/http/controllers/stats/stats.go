package stats

import (
	"fmt"
	"hr-server/internal/api/http/controllers/common"
	"hr-server/internal/api/http/controllers/stats/dto"
	"hr-server/internal/domain"
	"hr-server/internal/register"
	"hr-server/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type StatsController struct {
	userRepo    *repository.UserRepository
	channelRepo *repository.ChannelRepository
}

func NewStatsController(sr *register.StorageRegister) *StatsController {
	return &StatsController{
		userRepo:    sr.UserRepository(),
		channelRepo: sr.ChannelRepository(),
	}
}

// GetStats godoc
// @Summary Get overall statistics
// @Description Get overall statistics for the system
// @Tags Statistics
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetStatsResponse
// @Failure 500 {object} common.ErrorResponse
// @Security XAuthToken
// @Router /stats/all [get]
func (c *StatsController) GetStatsHandler(sr *register.StorageRegister) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// This would need to be implemented to get overall stats
		// For now, return a placeholder response
		response := dto.NewGetStatsResponse("Overall statistics not implemented yet")
		ctx.JSON(http.StatusOK, response)
	}
}

// GetChannelsStats godoc
// @Summary Get all channels statistics
// @Description Get statistics for all channels
// @Tags Statistics
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetChannelsStatsResponse
// @Failure 500 {object} common.ErrorResponse
// @Security XAuthToken
// @Router /stats/channels [get]
func (c *StatsController) GetChannelsStatsHandler(sr *register.StorageRegister) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// This would need to be implemented to get channel stats
		// For now, return a placeholder response
		response := dto.NewGetChannelsStatsResponse("Channel statistics not implemented yet")
		ctx.JSON(http.StatusOK, response)
	}
}

// GetChannelStats godoc
// @Summary Get channel statistics
// @Description Get statistics for a specific channel code
// @Tags Statistics
// @Accept json
// @Produce json
// @Param code path string true "Channel code"
// @Success 200 {object} domain.ChannelStats
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security XAuthToken
// @Router /stats/channel/{code} [get]
func (c *StatsController) GetChannelStatsHandler(sr *register.StorageRegister) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.Param("code")
		if code == "" {
			ctx.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Channel code is required"})
			return
		}

		// Get channel by code
		channel, err := c.channelRepo.GetByCode(code)
		if err != nil {
			logrus.Error("error while get channel by code: ", err)
			ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: fmt.Sprintf("failed to get channel by code '%s': %v", code, err)})
			return
		}

		if channel == nil {
			ctx.JSON(http.StatusNotFound, common.ErrorResponse{Error: "Channel code not found"})
			return
		}

		// Get users by channel
		users, err := c.userRepo.GetByChannel(channel.ID)
		if err != nil {
			logrus.Error("error while get users by channel: ", err)
			ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: fmt.Sprintf("failed to get users by channel %d: %v", channel.ID, err)})
			return
		}

		stats := domain.ChannelStats{
			ChannelID:   channel.ID,
			ChannelName: channel.Name,
			Code:        channel.Code,
			UserCount:   len(users),
		}

		response := dto.NewGetChannelStatsResponse(&stats)
		ctx.JSON(http.StatusOK, response)
	}
}
