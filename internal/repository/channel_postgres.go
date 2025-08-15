package repository

import (
	"errors"
	"fmt"
	"hr-server/internal/domain"
	"time"

	"gorm.io/gorm"
)

const CHANNELS_TABLE_NAME = "channels"

type PostgresChannel struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:255"`
	Code      string `gorm:"size:50;uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPostgresChannel(channel *domain.Channel) PostgresChannel {
	return PostgresChannel{
		ID:   channel.ID,
		Name: channel.Name,
		Code: channel.Code,
	}
}

func (pc PostgresChannel) TableName() string {
	return CHANNELS_TABLE_NAME
}

func (pc PostgresChannel) ToDomain() *domain.Channel {
	return &domain.Channel{
		ID:        pc.ID,
		Name:      pc.Name,
		Code:      pc.Code,
		CreatedAt: pc.CreatedAt,
		UpdatedAt: pc.UpdatedAt,
	}
}

type ChannelRepository struct {
	db *gorm.DB
}

func NewChannelRepository(db *gorm.DB) *ChannelRepository {
	db.AutoMigrate(PostgresChannel{})
	return &ChannelRepository{db}
}

func (r *ChannelRepository) Create(name, code string) (*domain.Channel, error) {
	channel := &domain.Channel{
		Name: name,
		Code: code,
	}

	postgresChannel := NewPostgresChannel(channel)
	if err := r.db.Table(CHANNELS_TABLE_NAME).Create(&postgresChannel).Error; err != nil {
		return nil, fmt.Errorf("failed to create channel with name '%s' and code '%s': %w", name, code, err)
	}

	return postgresChannel.ToDomain(), nil
}

func (r *ChannelRepository) GetByCode(code string) (*domain.Channel, error) {
	var postgresChannel PostgresChannel

	if err := r.db.Table(CHANNELS_TABLE_NAME).First(&postgresChannel, "code = ?", code).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get channel by code '%s': %w", code, err)
	}

	return postgresChannel.ToDomain(), nil
}

func (r *ChannelRepository) GetByID(id int) (*domain.Channel, error) {
	var postgresChannel PostgresChannel

	if err := r.db.Table(CHANNELS_TABLE_NAME).First(&postgresChannel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get channel by ID %d: %w", id, err)
	}

	return postgresChannel.ToDomain(), nil
}

func (r *ChannelRepository) GetAll() ([]*domain.Channel, error) {
	var postgresChannels []PostgresChannel

	if err := r.db.Table(CHANNELS_TABLE_NAME).Find(&postgresChannels).Error; err != nil {
		return nil, fmt.Errorf("failed to get all channels: %w", err)
	}

	var channels []*domain.Channel
	for _, pc := range postgresChannels {
		channels = append(channels, pc.ToDomain())
	}

	return channels, nil
}
