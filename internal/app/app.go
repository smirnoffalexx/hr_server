package app

import (
	"context"
	"errors"
	"hr-server/config"
	"hr-server/internal/api/http/routing"
	"hr-server/internal/background"
	"hr-server/internal/register"
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
		return err
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	sr, err := register.NewStorageRegister(cfg)
	if err != nil {
		return err
	}

	logrus.Info("Service started")

	var wg sync.WaitGroup
	wg.Add(1)

	tgBot, err := background.NewTelegramBot(sr)
	if err != nil {
		return err
	}
	go tgBot.Run(ctx, &wg)

	router := gin.New()
	routing.SetGinMiddlewares(router)
	routing.SetRouterHandler(router, sr)

	server := &http.Server{
		Addr:    ":" + cfg.Http.Port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("http server error: %v", err)
		}
	}()

	<-ctx.Done()
	logrus.Warn("Shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logrus.Errorf("server shutdown error: %v", err)
	}

	wg.Wait()

	logrus.Info("Service gracefully stopped")

	return nil
}
