package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/BlackRRR/notion-control/config"
	"github.com/BlackRRR/notion-control/model"
	"github.com/BlackRRR/notion-control/msgs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var Len int

func StartListeningRequests(bot *tgbotapi.BotAPI, handlers *Handlers) {
	var count int64
	for range time.Tick(time.Second * 15) {
		url := config.Url

		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			log.Println(err)
		}

		req.Header.Add("Accept", "application/json")
		req.Header.Add("Notion-Version", "2022-02-22")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+config.Secret)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
		}

		CheckResBodyAndSend(bot, res, handlers)

		InitHandler()
		count++
		fmt.Println("number of req", count)
		model.TotalCountReq.WithLabelValues("notion-req").Inc()
	}
}

func CheckResBodyAndSend(bot *tgbotapi.BotAPI, res *http.Response, handlers *Handlers) {
	var Text string

	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &handlers.Results)
	if err != nil {
		log.Println(err)
	}

	if Len == len(body) || Len > len(body) {
		return
	}

	Len = len(body)

	Text = CriticalStatus(handlers, Text)

	Text = NoStatus(handlers, Text)

	if Text != "" {
		for _, value := range model.Admins {
			msg := msgs.CreateNewTGMessage(Text, value)
			msgs.BotSendMsg(bot, msg)
			model.TotalCountSend.WithLabelValues("notion").Inc()
		}
	}

	Text = ""
}

func NewMsgForATask(i int, handlers *Handlers, status string) string {
	var Text string
	Text = "New Task was added \n"

	Text = Text + "Status = " + status + "\n"

	Text = Text + "URL = " + handlers.Results.Page[i].Url

	return Text
}

func CheckTime(i int, handlers *Handlers) time.Time {
	var parse time.Time
	if handlers.Results.Page[i].CreatedTime != "" {
		parser, err := time.Parse(time.RFC3339, handlers.Results.Page[i].CreatedTime)
		if err != nil {
			log.Println(err)
		}

		parse = parser
	}

	return parse
}

func CriticalStatus(handlers *Handlers, Text string) string {
	for i := range handlers.Results.Page {
		parse := CheckTime(i, handlers)
		if handlers.Results.Page[i].Properties.Status.Select != nil {
			if handlers.Results.Page[i].Properties.Status.Select.Name == "Critical" && parse.Minute() == time.Now().Minute() || parse.Minute() == time.Now().Minute()-1 && time.Now().Day() == parse.Day() {
				Text = NewMsgForATask(i, handlers, "Critical")
				return Text
			}
		}
	}

	return Text
}

func NoStatus(handlers *Handlers, Text string) string {
	parse := CheckTime(0, handlers)
	if parse.Minute() == time.Now().Minute() || parse.Minute() == time.Now().Minute()-1 && time.Now().Day() == parse.Day() {
		if handlers.Results.Page[0].Properties.Status.Select == nil {
			return NewMsgForATask(0, handlers, "No Status")
		}
	}

	return Text
}
