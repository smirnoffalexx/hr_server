package service

import (
	"log"
	"time"

	"hr-server/internal/api/http/controllers/notification/dto"
	"hr-server/internal/domain"
	"hr-server/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	DefaultMessageInterval = 1 * time.Second
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
	Emoji    *string
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
func (s *NotificationService) SendNotification(req *dto.SendNotificationRequest) error {
	// Create jobs channel with reasonable capacity
	jobs := make(chan NotificationJob, DefaultBatchSize)

	// Start workers
	for w := 1; w <= DefaultWorkerCount; w++ {
		go s.worker(w, jobs)
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
					Message:  req.Message,
					ImageURL: req.ImageURL,
					Emoji:    req.Emoji,
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
func (s *NotificationService) worker(id int, jobs <-chan NotificationJob) {

	for job := range jobs {
		log.Printf("[Worker %d] Sending notification to user %d (@%s)",
			id, job.User.TelegramID, job.User.Username)

		// Add emoji to message if provided
		message := job.Message
		if job.Emoji != nil {
			message = *job.Emoji + " " + message
		}

		var err error
		if job.ImageURL != nil && *job.ImageURL != "" {
			photo := tgbotapi.NewPhoto(job.User.TelegramID, tgbotapi.FileURL(*job.ImageURL))
			photo.Caption = message
			photo.ParseMode = "Markdown"
			err = s.telegramService.SendMessage(job.User.TelegramID, photo)
		} else {
			msg := tgbotapi.NewMessage(job.User.TelegramID, message)
			msg.ParseMode = "Markdown"
			err = s.telegramService.SendMessage(job.User.TelegramID, msg)
		}

		if err != nil {
			log.Printf("failed to send to user %d: %v", job.User.TelegramID, err)
		} else {
			log.Printf("successfully sent to user %d", job.User.TelegramID)
		}

		// Rate limiting to avoid hitting Telegram API limits
		time.Sleep(DefaultMessageInterval)
	}
}
