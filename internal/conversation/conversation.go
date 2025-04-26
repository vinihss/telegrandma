package conversation

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"sync"
	"telegrandma/internal/core"
	"time"
)

var chatChannels = make(map[int64]chan tgbotapi.Update)
var mu sync.Mutex
var wg sync.WaitGroup

const numWorkers = 5 // Número de goroutines simultâneas

type BotSettings struct {
	Token         string
	Debug         bool
	UpdateOffset  int
	UpdateTimeout int
}

type Chat struct {
	ID         int
	UserName   string
	LastUpdate time.Time
}

func InitializeBot() {
	core.LogInfo("InitializeBot")
	settings := BotSettings{
		Token:         os.Getenv("TELEGRAM_API_TOKEN"),
		Debug:         false,
		UpdateOffset:  0,
		UpdateTimeout: 60,
	}

	bot, err := tgbotapi.NewBotAPI(settings.Token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = settings.Debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(settings.UpdateOffset)
	updateConfig.Timeout = settings.UpdateTimeout

	updates := bot.GetUpdatesChan(updateConfig)
	core.LogInfo("Listening for messages")

	for update := range updates {

		if update.Message == nil {
		}
		core.LogInfo("Received from: " + update.Message.From.UserName)
		go processMessage(bot, update)

	}
}

func processMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	text, err := core.ExecCommand(update.Message.Text)

	if err != nil {
		log.Printf("Erro ao executar o comando: %v", err)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
