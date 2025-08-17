package app

import (
	"context"
	"errors"
	"fmt"
	"hr-server/config"
	"hr-server/internal/api/http/routing"
	"hr-server/internal/infrastructure"
	"hr-server/internal/repository"
	"hr-server/internal/service"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Config) error {
	if err := InitLogger(cfg); err != nil {
		return fmt.Errorf("failed to init logger: %w", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, err := infrastructure.NewPostgresDB(cfg)
	if err != nil {
		return fmt.Errorf("failed to create postgres database: %w", err)
	}

	userRepository := repository.NewUserRepository(db)
	channelRepository := repository.NewChannelRepository(db)

	userService := service.NewUserService(userRepository)
	channelService := service.NewChannelService(channelRepository)

	var wg sync.WaitGroup
	wg.Add(1)

	telegramService, err := service.NewTelegramService(cfg, userService, channelService)
	if err != nil {
		return fmt.Errorf("failed to create telegram bot: %w", err)
	}

	notificationService := service.NewNotificationService(userRepository, telegramService)

	go telegramService.Run(ctx, &wg)

	router := gin.New()
	routing.SetGinMiddlewares(router)
	routing.SetRouterHandler(router, cfg, userService, channelService, notificationService)

	server := &http.Server{
		Addr:    ":" + cfg.Http.Port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("http server error: %v", err)
		}
	}()

	logrus.Info("service started")

	<-ctx.Done()
	logrus.Warn("Shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logrus.Errorf("server shutdown error: %v", err)
	}

	wg.Wait()

	logrus.Info("service gracefully stopped")

	return nil
}
