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

func (s *UserService) CreateUser(telegramID int64, username string, channelID *int) (*domain.User, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.GetByTelegramID(telegramID)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	if existingUser != nil {
		// Update channel if provided
		if channelID != nil {
			err = s.userRepo.UpdateChannel(telegramID, channelID)
			if err != nil {
				return nil, fmt.Errorf("failed to update user channel: %w", err)
			}

			// Get updated user
			existingUser, err = s.userRepo.GetByTelegramID(telegramID)
			if err != nil {
				return nil, fmt.Errorf("failed to get updated user: %w", err)
			}
		}

		return existingUser, nil
	}

	// Create new user
	user, err := s.userRepo.Create(telegramID, username, channelID)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
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
