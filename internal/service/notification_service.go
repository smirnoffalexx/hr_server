package service

import (
	"log"
	"time"

	"hr-server/internal/domain"
	"hr-server/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	DefaultMessageInterval = 100 * time.Millisecond
	DefaultBatchSize       = 20
	DefaultWorkerCount     = 5
)

type NotificationService struct {
	userRepo        *repository.UserRepository
	telegramService *TelegramService
}

type NotificationJob struct {
	User     *domain.User
	Message  string
	ImageURL *string
}

func NewNotificationService(
	userRepo *repository.UserRepository,
	telegramService *TelegramService,
) *NotificationService {
	return &NotificationService{
		userRepo:        userRepo,
		telegramService: telegramService,
	}
}

// SendNotification sends notification to ALL users without any exceptions or filters
func (s *NotificationService) SendNotification(data *domain.NotificationData) error {
	// Create jobs channel with reasonable capacity
	jobs := make(chan NotificationJob, DefaultBatchSize)

	// Start workers
	for w := 1; w <= DefaultWorkerCount; w++ {
		go s.worker(jobs)
	}

	// Start a goroutine to load users in batches and send jobs
	go func() {
		defer close(jobs)

		// Load ALL users in batches - no filters, no exceptions
		err := s.userRepo.GetAllInBatches(DefaultBatchSize, func(batch []*domain.User) error {
			for _, user := range batch {
				// Send to ALL users without any filters
				jobs <- NotificationJob{
					User:     user,
					Message:  data.Message,
					ImageURL: data.ImageURL,
				}
			}
			return nil
		})

		if err != nil {
			log.Printf("error loading users in batches: %v", err)
		}
	}()

	return nil
}

// worker processes notification jobs
func (s *NotificationService) worker(jobs <-chan NotificationJob) {
	for job := range jobs {
		var err error
		if job.ImageURL != nil && *job.ImageURL != "" {
			photo := tgbotapi.NewPhoto(job.User.TelegramID, tgbotapi.FileURL(*job.ImageURL))
			photo.Caption = job.Message
			photo.ParseMode = "Markdown"
			err = s.telegramService.SendMessage(job.User.TelegramID, photo)
		} else {
			msg := tgbotapi.NewMessage(job.User.TelegramID, job.Message)
			msg.ParseMode = "Markdown"
			err = s.telegramService.SendMessage(job.User.TelegramID, msg)
		}

		if err != nil {
			log.Printf("failed to send to user %d: %v", job.User.TelegramID, err)
		}

		// Rate limiting to avoid hitting Telegram API limits
		time.Sleep(DefaultMessageInterval)
	}
}
