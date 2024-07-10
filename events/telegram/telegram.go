package telegram

import "go_telegram_bot/clients/telegram"

type Proccessor struct {
	tg *telegram.Client
	offset int
	// storage
}

func New(client *telegram.Client, storage) *Proccessor {

}