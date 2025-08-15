package background

import (
	"context"
	"fmt"
	"hr-server/internal/register"
	"hr-server/internal/service"
	"strings"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type TelegramBot struct {
	bot            *tgbotapi.BotAPI
	userService    *service.UserService
	channelService *service.ChannelService
}

func NewTelegramBot(sr *register.StorageRegister) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(sr.Config().TgBotToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create telegram bot: %w", err)
	}

	return &TelegramBot{
		bot:            bot,
		userService:    sr.UserService(),
		channelService: sr.ChannelService(),
	}, nil
}

func (t *TelegramBot) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	logrus.Info("Telegram bot started")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := t.bot.GetUpdatesChan(u)

	for {
		select {
		case <-ctx.Done():
			logrus.Info("Telegram bot stopped")
			return
		case update := <-updates:
			if update.Message == nil {
				continue
			}

			if update.Message.IsCommand() {

				switch update.Message.Command() {
				case "start":
					if err := t.handleStartCommand(update.Message); err != nil {
						logrus.Errorf("failed to handle start command: %v", err)
					}
				default:
					continue
				}

				chatID := update.Message.Chat.ID
				msgText := "✅ Добро пожаловать!\n\n👋 Откройте Telegram mini app для продолжения."
				msg := tgbotapi.NewMessage(chatID, msgText)
				if _, err := t.bot.Send(msg); err != nil {
					logrus.Errorf("failed to send msg: %v", err)
				}
			}
		}
	}
}

func (t *TelegramBot) handleStartCommand(message *tgbotapi.Message) error {
	telegramID := message.From.ID
	username := message.From.UserName

	args := strings.Fields(message.Text)
	var channelID *int

	if len(args) > 1 {
		channelCode := args[1]

		if channelCode != "" {
			channel, err := t.channelService.GetChannelByCode(channelCode)
			if err != nil {
				return fmt.Errorf("failed to get channel by code %s: %v", channelCode, err)
			}

			if channel != nil {
				channelID = &channel.ID
			}
		}
	}

	if err := t.userService.CreateUser(telegramID, username, channelID); err != nil {
		return fmt.Errorf("failed to create user %d: %v", telegramID, err)
	}

	return nil
}
