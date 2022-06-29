package main

import (
	"github.com/BlackRRR/notion-control/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	bot := startBot()
	handler := handlers.InitHandler()
	go startPrometheusHandler()
	go handlers.StartListeningRequests(bot, handler)
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
