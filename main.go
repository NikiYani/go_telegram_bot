package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()

	// tgClient = telegram.New(token)

	// fetcher = fetcher.New()

	// proccessor = proccessor.New()

	// condumer.Start(fetcher, proccessor)

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
