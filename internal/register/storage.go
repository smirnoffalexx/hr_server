package register

import (
	"hr-server/config"
	"hr-server/internal/infrastructure"
	"hr-server/internal/repository"
	"hr-server/internal/service"

	"gorm.io/gorm"
)

type StorageRegister struct {
	cfg *config.Config

	db *gorm.DB

	userRepository *repository.UserRepository

	channelRepository *repository.ChannelRepository

	userService *service.UserService

	channelService *service.ChannelService

	notificationService *service.NotificationService
}

func NewStorageRegister(cfg *config.Config) (*StorageRegister, error) {
	sr := &StorageRegister{}
	sr.cfg = cfg

	db, err := infrastructure.NewPostgresDB(cfg)
	if err != nil {
		return nil, err
	}
	sr.db = db

	sr.userRepository = repository.NewUserRepository(db)
	sr.channelRepository = repository.NewChannelRepository(db)

	sr.userService = service.NewUserService(sr.userRepository)
	sr.channelService = service.NewChannelService(sr.channelRepository)
	sr.notificationService = service.NewNotificationService()

	return sr, nil
}

// UserRepository ...
func (r *StorageRegister) UserRepository() *repository.UserRepository {
	return r.userRepository
}

// ChannelRepository ...
func (r *StorageRegister) ChannelRepository() *repository.ChannelRepository {
	return r.channelRepository
}

// UserService ...
func (r *StorageRegister) UserService() *service.UserService {
	return r.userService
}

// ChannelService ...
func (r *StorageRegister) ChannelService() *service.ChannelService {
	return r.channelService
}

// NotificationService ...
func (r *StorageRegister) NotificationService() *service.NotificationService {
	return r.notificationService
}

// Config ...
func (r *StorageRegister) Config() *config.Config {
	return r.cfg
}

// DB ...
func (r *StorageRegister) DB() *gorm.DB {
	return r.db
}
