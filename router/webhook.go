package router

import (
	"context"
	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/db"
	cryptoFetcher "sayuri_crypto_bot/fetcher/crypto"
	"sayuri_crypto_bot/fortune/tarot"
	"sayuri_crypto_bot/model"
	"sayuri_crypto_bot/sender"
	"sayuri_crypto_bot/template"
	"strings"
	"time"
)

var commandFuncs map[string]func(ctx context.Context, params []string)

func initCommandFuncMap() {
	commandFuncs = make(map[string]func(ctx context.Context, params []string), 3)
	commandFuncs["/about"] = aboutCommand
	commandFuncs["/realtime"] = checkUserChatAvailable(realtimeCommand)
	commandFuncs["/tarot"] = checkUserAvailable(tarotCommand)
	commandFuncs["/sticker"] = sendSticker
}

func HandleCommands() {

	initCommandFuncMap()

	tgbot, err := bot.NewBotAPI(conf.GetConfig().Tgbot.Token)
	if err != nil {
		panic(err)
	}

	updateConfig := bot.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := tgbot.GetUpdatesChan(updateConfig)

	for update := range updates {
		ctx := context.Background()
		context.WithValue(ctx, conf.LOG_KEY_REQUEST_ID, uuid.NewString())
		go handleTgbotMessage(ctx, update.Message)
	}
}

func handleTgbotMessage(c context.Context, m *bot.Message) {
	log := conf.GetLog(c)
	log.Debugf("message: %+v\n", *m)
	user := m.From
	chat := m.Chat
	text := m.Text
	if len(text) == 0 {
		return
	}
	newCtx := context.WithValue(c, "tg_user", user.ID)
	newCtx = context.WithValue(newCtx, "tg_chat", chat.ID)
	// command
	if text[0] == '/' {
		sls := strings.Split(text, " ")
		cmd := getCommand(sls[0])
		f, ok := commandFuncs[cmd]
		if ok {
			go f(newCtx, sls)
		} else {
			log.Error("cannot find command: ", cmd)
		}
	}
}

func checkUserAvailable(f func(ctx context.Context, params []string)) func(ctx context.Context, params []string) {
	return func(ctx context.Context, params []string) {
		log := conf.GetLog(ctx)
		user, ok := ctx.Value("tg_user").(int64)
		if !ok {
			log.Error("cannot get telegram user id from context")
			return
		}
		chat, ok := ctx.Value("tg_chat").(int64)
		if !db.CheckUserAvailable(ctx, user) {
			msg, _ := template.GetString(template.TooOften, nil)
			if ok {
				sender.TgSendData(chat, msg)
			} else {
				sender.TgSendData(user, msg)
			}
			return
		}
		f(ctx, params)
	}
}

func checkChatAvailable(f func(ctx context.Context, params []string)) func(ctx context.Context, params []string) {
	return func(ctx context.Context, params []string) {
		log := conf.GetLog(ctx)
		chat, ok := ctx.Value("tg_chat").(int64)
		if !ok {
			log.Error("cannot get telegram chat id from context")
			return
		}
		if !db.CheckChatAvailable(ctx, chat) {
			msg, _ := template.GetString(template.TooOften, nil)
			sender.TgSendData(chat, msg)
			return
		}
		f(ctx, params)
	}
}

func checkUserChatAvailable(f func(ctx context.Context, params []string)) func(ctx context.Context, params []string) {
	return checkUserAvailable(checkChatAvailable(f))
}

func aboutCommand(ctx context.Context, params []string) {
	log := conf.GetLog(ctx)
	chat, ok := ctx.Value("tg_chat").(int64)
	if !ok {
		log.Error("cannot get telegram chat id from context")
		return
	}
	msg, _ := template.GetString(template.Aboutme, nil)
	sender.TgSendData(chat, msg)
}

func realtimeCommand(ctx context.Context, params []string) {
	log := conf.GetLog(ctx)
	user, ok := ctx.Value("tg_user").(int64)
	if !ok {
		log.Error("cannot get telegram user id from context")
		return
	}
	chat, ok := ctx.Value("tg_chat").(int64)
	if !ok {
		log.Error("cannot get telegram chat id from context")
		return
	}

	infoMsg, err := template.GetString(template.Realtime, nil)
	if err != nil {
		log.Error("error: ", err)
		return
	}
	if len(params) != 2 {
		sender.TgSendData(chat, infoMsg)
		return
	}

	// crypto
	item, err := db.GetCryptoItemById(ctx, params[1])
	if err != nil {
		log.Error("error: ", err)
	}
	if item != nil {
		items := []*model.GoodsItem{item}
		markets, err := cryptoFetcher.GetCryptoFetcher().GetValue(items)
		if err != nil {
			log.Info("error: ", err)
			return
		}
		output := model.Output{
			Datetime: time.Now().Format("2006-01-02 15:04:05"),
			Items:    markets,
		}
		msg, err := template.GetString(template.Crypto, output)
		if err != nil {
			log.Error("error: ", err)
			return
		} else {
			sender.TgSendData(chat, msg)
			db.AddChatPeriod(ctx, chat)
			db.AddUserPeriod(ctx, user)
			return
		}
	}
	sender.TgSendData(chat, infoMsg)

}

func getCommand(cmd string) string {
	sls := strings.Split(cmd, "@")
	return sls[0]
}

func tarotCommand(ctx context.Context, params []string) {
	log := conf.GetLog(ctx)
	user, ok := ctx.Value("tg_user").(int64)
	if !ok {
		log.Error("cannot get telegram user id from context")
		return
	}
	chat, ok := ctx.Value("tg_chat").(int64)
	if !ok {
		log.Error("cannot get telegram chat id from context")
		return
	}
	card := tarot.Draw()
	msg, err := template.GetString(template.Tarot, card)
	if err != nil {
		log.Error("error: ", err)
		return
	}
	sender.TgSendData(chat, msg)
	db.AddChatPeriod(ctx, chat)
	db.AddUserPeriod(ctx, user)
}

func sendSticker(ctx context.Context, params []string) {
	log := conf.GetLog(ctx)
	chat, ok := ctx.Value("tg_chat").(int64)
	if !ok {
		log.Error("cannot get telegram chat id from context")
		return
	}
	sender.TgSendSticker(chat)
}
