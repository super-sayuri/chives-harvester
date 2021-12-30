package sender

import (
	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sayuri_crypto_bot/conf"
	"strconv"
)

func TgSendData(userId int64, msg string) error {
	tgbot, err := bot.NewBotAPI(conf.GetConfig().Tgbot.Token)
	if err != nil {
		return err
	}
	c := bot.MessageConfig{

		BaseChat: bot.BaseChat{
			ChatID: userId,
		},
		Text: msg,
	}
	_, err = tgbot.Send(c)
	if err != nil {
		return err
	}
	return nil
}

func TgStartMessage(ownerId string) error {
	id, err := strconv.Atoi(ownerId)
	if err != nil {
		return err
	}
	tgbot, err := bot.NewBotAPI(conf.GetConfig().Tgbot.Token)
	if err != nil {
		return err
	}
	c := bot.MessageConfig{

		BaseChat: bot.BaseChat{
			ChatID: int64(id),
		},
		Text: "天灾再起! " + conf.GetConfig().Common.Name,
	}
	_, err = tgbot.Send(c)
	if err != nil {
		return err
	}
	return nil

}
