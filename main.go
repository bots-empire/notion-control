package main

import (
	"github.com/BlackRRR/notion-control/handlers"
	"github.com/BlackRRR/notion-control/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	bot, updates := startBot()
	model.DownloadAdminSettings()
	handler := handlers.InitHandler()
	go startPrometheusHandler()
	go handlers.StartListeningRequests(bot, handler)
	startHandlers(bot, updates)
	sig := <-subscribeToSystemSignals()

	log.Printf("shutdown all process on '%s' system signal\n", sig.String())
}

func startBot() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	file, err := os.ReadFile("./config/token.txt")
	if err != nil {
		return nil, nil
	}

	bot, err := tgbotapi.NewBotAPI(string(file))
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)

	channel := bot.GetUpdatesChan(u)

	globalBot := handlers.BotInit(channel)

	log.Println("The bot is running")

	return bot, globalBot.Update
}

func startHandlers(bot *tgbotapi.BotAPI, update tgbotapi.UpdatesChannel) {
	//wg := new(sync.WaitGroup)
	go func() {
		handlers.ActionWithUpdates(bot, update)
	}()

	//service.Msgs.SendNotificationToDeveloper("Bot are restart", false)
	log.Println("All handlers are running")
}

func startPrometheusHandler() {
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Metrics can be read from %s port", "7011")
	metricErr := http.ListenAndServe(":7011", nil)
	if metricErr != nil {
		log.Fatalf("metrics stoped by metricErr: %s\n", metricErr.Error())
	}
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
