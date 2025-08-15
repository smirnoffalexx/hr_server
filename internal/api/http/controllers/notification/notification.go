package notification

import (
	"fmt"
	"hr-server/internal/api/http/controllers/common"
	"hr-server/internal/api/http/controllers/notification/dto"
	"hr-server/internal/domain"
	"hr-server/internal/register"
	"hr-server/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type NotificationController struct {
	notificationService *service.NotificationService
}

func NewNotificationController(sr *register.StorageRegister) *NotificationController {
	return &NotificationController{sr.NotificationService()}
}

// SendNotification godoc
// @Summary Send notification to all users
// @Description Send a notification message to all users
// @param X-Auth-Token header string true "X-Auth-Token"
// @Tags Notifications
// @Accept json
// @Produce json
// @Param request body dto.SendNotificationRequest true "Send notification request"
// @Success 200 {object} common.SuccessResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Security XAuthToken
// @Router /api/notifications [post]
func (c *NotificationController) SendNotificationHandler(sr *register.StorageRegister) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := dto.NewSendNotificationRequest()
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

		domainReq := &domain.SendNotificationRequest{
			Message:  req.Message,
			ImageURL: req.ImageURL,
			Emoji:    req.Emoji,
		}

		err := c.notificationService.SendNotification(domainReq)
		if err != nil {
			logrus.Error("error while send notification: ", err)
			ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: fmt.Sprintf("failed to send notification: %v", err)})
			return
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse{})
	}
}
