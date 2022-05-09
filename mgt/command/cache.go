package command

import (
	"context"
	"encoding/json"
	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/db"
)

func updateCache(ctx context.Context) ([]bot.BotCommand, error) {
	tgbot, err := bot.NewBotAPI(conf.GetConfig().Tgbot.Token)
	if err != nil {
		return nil, err
	}
	commands, err := tgbot.GetMyCommands()
	if err != nil {
		return nil, err
	}
	go func(ctx context.Context, commands []bot.BotCommand) {
		log := conf.GetLog(ctx)
		data, err := json.Marshal(commands)
		if err != nil {
			log.Error("cannot parse tgbot commands, ", err)
			return
		}
		err = db.SaveTgbotCommands(ctx, string(data))
		if err != nil {
			log.Error("cannot cache tgbot commands, ", err)
			return
		}
	}(ctx, commands)
	return commands, nil
}

func readFromCache(ctx context.Context) ([]bot.BotCommand, error) {
	data, err := db.GetTgbotCommands(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]bot.BotCommand, 0)
	err = json.Unmarshal([]byte(data), &res)
	return res, nil
}

func read(ctx context.Context) ([]bot.BotCommand, error) {
	log := conf.GetLog(ctx)
	res, err := readFromCache(ctx)
	if err != nil {
		log.Error("Cannot read tgbot commands from cache: ", err)
		return nil, err
	}
	if len(res) != 0 {
		return res, nil
	}
	log.Info("No tgbot commands in its cache")
	res, err = updateCache(ctx)
	if err != nil {
		log.Error("Cannot update tgbot commands in cache: ", err)
		return nil, err
	}
	return res, nil
}
