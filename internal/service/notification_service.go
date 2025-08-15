package service

import "hr-server/internal/domain"

type NotificationService struct {
}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) SendNotification(req *domain.SendNotificationRequest) error {
	return nil
}
