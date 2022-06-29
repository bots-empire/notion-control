package main

import (
	"github.com/BlackRRR/notion-control/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	bot := startBot()
	handler := handlers.InitHandler()
	StartListenNotionReq(bot, handler)

	sig := <-subscribeToSystemSignals()

	log.Printf("shutdown all process on '%s' system signal\n", sig.String())
}

func startBot() *tgbotapi.BotAPI {
	file, err := os.ReadFile("./config/token.txt")
	if err != nil {
		return nil
	}

	bot, err := tgbotapi.NewBotAPI(string(file))
	if err != nil {
		log.Panic(err)
	}

	log.Println("The bot is running")

	return bot
}

func StartListenNotionReq(bot *tgbotapi.BotAPI, handler *handlers.Handlers) {
	go handlers.StartListeningRequests(bot, handler)
}

func subscribeToSystemSignals() chan os.Signal {
	ch := make(chan os.Signal, 10)
	signal.Notify(ch,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGHUP,
	)
	return ch
}
