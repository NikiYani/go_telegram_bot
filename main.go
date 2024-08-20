package main

import (
	"flag"
	event_consumer "github.com/go_telegram_bot/consumer/event-consumer"
	"log"

	tgClient "github.com/RSheremeta/read-adviser-bot/clients/telegram"
	"github.com/RSheremeta/read-adviser-bot/consumer/event-consumer"
	"github.com/RSheremeta/read-adviser-bot/events/telegram"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("Token is not specified")
	}

	return *token
}
