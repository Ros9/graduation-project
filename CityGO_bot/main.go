package main

import (
	"flag"
	"log"

	tgClient "graduation-project/CityGO_bot/clients/telegram"
	event_consumer "graduation-project/CityGO_bot/consumer/event-consumer"
	"graduation-project/CityGO_bot/events/telegram"
	commandshistory "graduation-project/CityGO_bot/lib/commandsHistory"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

var commandHistory []commandshistory.CommandHistoryItem

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
	)

	//commandHistory = append(commandHistory, models.CommandHistoryItem{ChatId: 1, Text: "TEST"})

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(&commandHistory); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
