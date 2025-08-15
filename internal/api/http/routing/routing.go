package routing

import (
	"bytes"
	"hr-server/internal/api/http/controllers/channel"
	"hr-server/internal/api/http/controllers/notification"
	"hr-server/internal/api/http/controllers/stats"
	"hr-server/internal/api/http/controllers/user"
	_ "hr-server/internal/api/http/docs"
	"hr-server/internal/api/http/middleware"
	"hr-server/internal/register"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var ignoredPrefixPaths = []string{
	"/api/health",
	"/api/swagger",
}

func GinLogrusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		var requestBody []byte
		if c.Request.Body != nil {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				requestBody = bodyBytes
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // recover request body
			}
		}

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()
		path := c.Request.URL.Path

		for _, prefix := range ignoredPrefixPaths {
			if strings.HasPrefix(path, prefix) {
				return
			}
		}

		logrus.WithFields(logrus.Fields{
			"status":    status,
			"method":    c.Request.Method,
			"path":      path,
			"query":     c.Request.URL.RawQuery,
			"body":      string(requestBody),
			"latency":   duration,
			"client_ip": c.ClientIP(),
		}).Info("incoming request")
	}
}

func SetGinMiddlewares(router *gin.Engine) {
	router.Use(GinLogrusMiddleware())
	router.Use(gin.Recovery()) // recovery middleware
}

// @title HR Server API
// @version 1.0.0
// @description This is the Swagger documentation for the HR Server service.
// @BasePath /api
// @securityDefinitions.apikey XAuthToken
// @in header
// @name X-Auth-Token
// @description Enter your authentication token
func SetRouterHandler(router *gin.Engine, sr *register.StorageRegister) {
	apiGroup := router.Group("/api")

	apiGroup.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "pong",
		})
	})

	apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup.Use(middleware.AuthTokenMiddleware(sr.Config().AuthToken))

	// User routes
	userGroup := apiGroup.Group("/users")
	userController := user.NewUserController(sr)
	userGroup.GET("/", userController.GetUsersHandler(sr))

	// Channel routes
	channelGroup := apiGroup.Group("/channels")
	channelController := channel.NewChannelController(sr)
	channelGroup.POST("/generate", channelController.GenerateChannelHandler(sr))
	channelGroup.GET("/:code", channelController.GetChannelByCodeHandler(sr))
	channelGroup.POST("/bulk", channelController.GenerateBulkChannelHandler(sr))
	channelGroup.GET("/all", channelController.GetChannelsHandler(sr))

	// Notification routes
	notificationGroup := apiGroup.Group("/notifications")
	notificationController := notification.NewNotificationController(sr)
	notificationGroup.POST("/", notificationController.SendNotificationHandler(sr))

	// Stats routes
	statsGroup := apiGroup.Group("/stats")
	statsController := stats.NewStatsController(sr)
	statsGroup.GET("/channels", statsController.GetChannelsStatsHandler(sr))
	statsGroup.GET("/channel/:code", statsController.GetChannelStatsHandler(sr))
	statsGroup.GET("/all", statsController.GetStatsHandler(sr))
}
