package user

import (
	"encoding/csv"
	"fmt"
	"hr-server/internal/api/http/controllers/common"
	"hr-server/internal/api/http/controllers/user/dto"
	"hr-server/internal/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService}
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all registered users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetUsersResponse
// @Failure 500 {object} common.ErrorResponse
// @Security XAuthToken
// @Router /users [get]
func (c *UserController) GetUsersHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := c.userService.GetAllUsersWithChannel()
		if err != nil {
			logrus.Error("error while get all users: ", err)
			ctx.JSON(
				http.StatusInternalServerError,
				common.ErrorResponse{Error: fmt.Sprintf("failed to get all users: %v", err)},
			)
			return
		}

		response := dto.NewGetUsersResponse(users)
		ctx.JSON(http.StatusOK, response)
	}
}

// ExportUsers godoc
// @Summary Export all users to CSV
// @Description Export all registered users with channel information to CSV format
// @Tags Users
// @Accept json
// @Produce text/csv
// @Success 200 {file} file CSV file with users data
// @Failure 500 {object} common.ErrorResponse
// @Security XAuthToken
// @Router /users/export [get]
func (c *UserController) ExportUsersHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := c.userService.GetAllUsersWithChannel()
		if err != nil {
			logrus.Error("error while get all users for export: ", err)
			ctx.JSON(
				http.StatusInternalServerError,
				common.ErrorResponse{Error: fmt.Sprintf("failed to get all users: %v", err)},
			)
			return
		}

		// Set CSV headers
		ctx.Header("Content-Type", "text/csv")
		ctx.Header(
			"Content-Disposition",
			"attachment; filename=users_export_"+time.Now().Format("2006-01-02")+".csv",
		)

		// Create CSV writer
		writer := csv.NewWriter(ctx.Writer)
		defer writer.Flush()

		// Write CSV header
		header := []string{
			"ID",
			"Telegram ID",
			"Username",
			"Channel ID",
			"Channel Name",
			"Bot Start Link",
			"Created At",
			"Updated At",
		}
		if err := writer.Write(header); err != nil {
			logrus.Error("error while writing CSV header: ", err)
			ctx.JSON(
				http.StatusInternalServerError,
				common.ErrorResponse{Error: "failed to generate CSV"},
			)
			return
		}

		// Write user data
		for _, user := range users {
			channelID := ""
			if user.ChannelID != nil {
				channelID = strconv.Itoa(*user.ChannelID)
			}

			channelName := ""
			if user.ChannelName != nil {
				channelName = *user.ChannelName
			}

			row := []string{
				strconv.Itoa(user.ID),
				strconv.FormatInt(user.TelegramID, 10),
				user.Username,
				channelID,
				channelName,
				user.CreatedAt.Format("2006-01-02 15:04:05"),
				user.UpdatedAt.Format("2006-01-02 15:04:05"),
			}

			if err := writer.Write(row); err != nil {
				logrus.Error("error while writing CSV row: ", err)
				ctx.JSON(
					http.StatusInternalServerError,
					common.ErrorResponse{Error: "failed to generate CSV"},
				)
				return
			}
		}
	}
}
