package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"hr-server/internal/domain"
	"hr-server/internal/repository"
)

type ChannelService struct {
	channelRepo *repository.ChannelRepository
}

func NewChannelService(channelRepo *repository.ChannelRepository) *ChannelService {
	return &ChannelService{
		channelRepo: channelRepo,
	}
}

func (s *ChannelService) GenerateChannel(channelName string) (*domain.Channel, error) {
	// Generate unique channel code
	code, err := s.generateUniqueCode()
	if err != nil {
		return nil, fmt.Errorf("failed to generate unique code: %w", err)
	}

	// Create channel with channel code in database
	channel, err := s.channelRepo.Create(channelName, code)
	if err != nil {
		return nil, fmt.Errorf("failed to create channel: %w", err)
	}

	// Log the channel creation
	fmt.Printf("🎯 New Channel Code: %s\n\nUse this code to track your visits!\n\nBot: @YourBotName\nStart command: /start %s", code, code)

	return channel, nil
}

func (s *ChannelService) GenerateBulkChannel(channelNames []string) ([]*domain.Channel, error) {
	var channels []*domain.Channel

	for i, channelName := range channelNames {
		channel, err := s.GenerateChannel(channelName)
		if err != nil {
			return nil, fmt.Errorf("failed to generate channel code for '%s' at index %d: %w", channelName, i+1, err)
		}
		channels = append(channels, channel)
	}

	return channels, nil
}

func (s *ChannelService) GetChannelByCode(code string) (*domain.Channel, error) {
	return s.channelRepo.GetByCode(code)
}

func (s *ChannelService) GetAll() ([]*domain.Channel, error) {
	return s.channelRepo.GetAll()
}

func (s *ChannelService) GetChannelByID(id int) (*domain.Channel, error) {
	return s.channelRepo.GetByID(id)
}

func (s *ChannelService) generateUniqueCode() (string, error) {
	// generate 16-character hex code
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	return hex.EncodeToString(bytes), nil
}
