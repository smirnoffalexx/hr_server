package service

import (
	"fmt"
	"hr-server/internal/domain"
	"hr-server/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(
	telegramID int64,
	username string,
	channelID *int,
) error {
	existingUser, err := s.userRepo.GetByTelegramID(telegramID)
	if err != nil {
		return fmt.Errorf("failed to check existing user: %w", err)
	}

	if existingUser != nil {
		return nil
	}

	err = s.userRepo.Create(telegramID, username, channelID)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *UserService) GetUser(telegramID int64) (*domain.User, error) {
	return s.userRepo.GetByTelegramID(telegramID)
}

func (s *UserService) GetAllUsers() ([]*domain.User, error) {
	return s.userRepo.GetAll()
}

func (s *UserService) GetUsersByChannel(channelID int) ([]*domain.User, error) {
	return s.userRepo.GetByChannel(channelID)
}

func (s *UserService) GetAllUsersWithChannel() ([]*domain.UserWithChannel, error) {
	return s.userRepo.GetAllWithChannel()
}
