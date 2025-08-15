package main

import (
	"fmt"
	"hr-server/config"
	"hr-server/internal/app"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("failed to load config: %w", err)
	}

	if err := app.Run(cfg); err != nil {
		logrus.Error(err)
		return fmt.Errorf("failed to run app: %w", err)
	}

	return nil
}
