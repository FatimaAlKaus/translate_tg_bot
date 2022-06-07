package bot

import (
	"log"
	"translateBot/internal/config"
	"translateBot/internal/translator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	updateTimeOut int
	debug         bool
	bot           *tgbotapi.BotAPI
	updateConfig  *tgbotapi.UpdateConfig
	translator    translator.Translator
}

func (bot *Bot) Listen(logger *log.Logger) {
	updates := bot.bot.GetUpdatesChan(*bot.updateConfig)

	for update := range updates {
		if update.Message != nil { // If we got a message
			go func() {
				logger.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				resp, err := bot.translator.Translate("en", "ru", update.Message.Text)
				if err != nil {
					logger.Fatal(err)
					return
				}
				logger.Printf(*resp)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, *resp)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.bot.Send(msg)
			}()
		}
	}
}
func New(cfg config.Config, updateTimeOut int, debug bool) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.API_KEY)
	if err != nil {
		return nil, err
	}
	bot.Debug = debug
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return &Bot{
		bot:           bot,
		debug:         debug,
		updateTimeOut: updateTimeOut,
		updateConfig:  &u,
		translator:    translator.NewGoogleTranslator(cfg.RAPID_API_KEY),
	}, nil
}
