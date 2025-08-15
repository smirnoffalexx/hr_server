package repository

import (
	"errors"
	"fmt"
	"hr-server/internal/domain"
	"time"

	"gorm.io/gorm"
)

const USERS_TABLE_NAME = "users"

type PostgresUser struct {
	ID         int    `gorm:"primaryKey;autoIncrement"`
	TelegramID int64  `gorm:"uniqueIndex"`
	Username   string `gorm:"size:255"`
	ChannelID  *int   `gorm:"index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PostgresUserWithChannel struct {
	PostgresUser
	Channel PostgresChannel `gorm:"foreignKey:ChannelID"`
}

func NewPostgresUser(user *domain.User) PostgresUser {
	return PostgresUser{
		ID:         user.ID,
		TelegramID: user.TelegramID,
		Username:   user.Username,
		ChannelID:  user.ChannelID,
	}
}

func (pu PostgresUser) TableName() string {
	return USERS_TABLE_NAME
}

func (pu PostgresUser) ToDomain() *domain.User {
	return &domain.User{
		ID:         pu.ID,
		TelegramID: pu.TelegramID,
		Username:   pu.Username,
		ChannelID:  pu.ChannelID,
		CreatedAt:  pu.CreatedAt,
		UpdatedAt:  pu.UpdatedAt,
	}
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	if err := db.AutoMigrate(PostgresUser{}); err != nil {
		panic(err)
	}

	return &UserRepository{db}
}

func (r *UserRepository) Create(telegramID int64, username string, channelID *int) error {
	user := &domain.User{
		TelegramID: telegramID,
		Username:   username,
		ChannelID:  channelID,
	}

	postgresUser := NewPostgresUser(user)
	if err := r.db.Table(USERS_TABLE_NAME).Create(&postgresUser).Error; err != nil {
		return fmt.Errorf("failed to create user in database: %w", err)
	}

	return nil
}

func (r *UserRepository) GetByTelegramID(telegramID int64) (*domain.User, error) {
	var postgresUser PostgresUser

	if err := r.db.Table(USERS_TABLE_NAME).First(&postgresUser, "telegram_id = ?", telegramID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by telegram ID %d: %w", telegramID, err)
	}

	return postgresUser.ToDomain(), nil
}

func (r *UserRepository) GetAll() ([]*domain.User, error) {
	var postgresUsers []PostgresUser

	if err := r.db.Table(USERS_TABLE_NAME).Find(&postgresUsers).Error; err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	var users []*domain.User
	for _, pu := range postgresUsers {
		users = append(users, pu.ToDomain())
	}

	return users, nil
}

func (r *UserRepository) GetByChannel(channelID int) ([]*domain.User, error) {
	var postgresUsers []PostgresUser

	if err := r.db.Table(USERS_TABLE_NAME).Where("channel_id = ?", channelID).Find(&postgresUsers).Error; err != nil {
		return nil, fmt.Errorf("failed to get users by channel %d: %w", channelID, err)
	}

	var users []*domain.User
	for _, pu := range postgresUsers {
		users = append(users, pu.ToDomain())
	}

	return users, nil
}
