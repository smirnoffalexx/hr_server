package main

import (
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
		return err
	}

	if err := app.Run(cfg); err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
