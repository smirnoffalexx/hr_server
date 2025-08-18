package user

import (
	"fmt"
	"hr-server/internal/api/http/controllers/common"
	"hr-server/internal/api/http/controllers/user/dto"
	"hr-server/internal/service"
	"net/http"

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
