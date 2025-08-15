package infrastructure

import (
	"fmt"
	"hr-server/config"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Postgres.HOST,
		cfg.Postgres.USER,
		cfg.Postgres.PASSWORD,
		cfg.Postgres.DB,
		cfg.Postgres.PORT,
		"disable",
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %w", err)
	}

	var version string
	if err := db.Raw("select version()").Scan(&version).Error; err != nil {
		return nil, fmt.Errorf("failed to check postgres version: %w", err)
	}

	logrus.Info("PostgreSQL version:", version)

	if cfg.Logger.LOGLVL == "debug" {
		db = db.Debug()
	}

	return db, nil
}
