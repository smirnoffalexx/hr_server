package background

import (
	"context"
	"encoding/base64"
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
				chatID := update.Message.Chat.ID
				msgText := ""
				var err error

				switch update.Message.Command() {
				case "start":
					msgText, err = t.handleStartCommand(update.Message)
					if err != nil {
						logrus.Errorf("failed to handle start command: %v", err)
					}
				default:
					continue
				}

				msg := tgbotapi.NewMessage(chatID, msgText)
				_, err = t.bot.Send(msg)
				if err != nil {
					logrus.Errorf("failed to send msg: %v", err)
				}
			}
		}
	}
}

func (t *TelegramBot) handleStartCommand(message *tgbotapi.Message) (string, error) {
	telegramID := message.From.ID
	username := message.From.UserName

	args := strings.Fields(message.Text)
	var channelCode *string
	var channelID *int

	if len(args) > 1 {
		param := args[1]

		if strings.Contains(param, "=") {
			if decoded, err := base64.URLEncoding.DecodeString(param); err == nil {
				logrus.Infof("Parsed deep link params: %+v", decoded)
			}
		} else {
			channelCode = &param
		}

		if channelCode != nil {
			channel, err := t.channelService.GetChannelByCode(*channelCode)
			if err != nil {
				logrus.Errorf("failed to get channel by code %s: %v", *channelCode, err)
				return "", fmt.Errorf("failed to get channel by code %s: %v", *channelCode, err)
			}

			if channel != nil {
				channelID = &channel.ID
			}
		}
	}

	_, err := t.userService.CreateUser(telegramID, username, channelID)
	if err != nil {
		logrus.Errorf("failed to create/update user %d: %v", telegramID, err)
		return "", fmt.Errorf("failed to create/update user %d: %v", telegramID, err)
	}

	code := "undefined"
	if channelCode != nil {
		code = *channelCode
	}

	msgText := fmt.Sprintf(
		"✅ Добро пожаловать, %s!\n\n👋 Вы успешно зарегистрированы в системе.\n\n💡 Для подключения к каналу используйте команду:\n/start %s",
		username,
		code,
	)

	logrus.Infof("User %s (ID: %d) registered successfully", username, telegramID)
	return msgText, nil
}
