package router

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/db"
	"sayuri_crypto_bot/fetcher"
	"sayuri_crypto_bot/model"
	"sayuri_crypto_bot/sender"
	"sayuri_crypto_bot/template"
	"strings"
	"time"
)

const API_WEBHOOK = "api_webhook"

var commandFuncs map[string]func(ctx context.Context, params []string)

func initCommandFuncMap() error {
	commandFuncs = make(map[string]func(ctx context.Context, params []string), 0)
	commandFuncs["/about"] = aboutCommand
	commandFuncs["/realtime"] = checkUserChatAvailable(realtimeCommand)

	return nil
}

func webhookRouter(r *gin.RouterGroup) {
	r.POST(routerMap[API_WEBHOOK], func(c *gin.Context) {
		log := conf.GetLog(c)
		defer NormalResponse(c, nil)
		tgbot, _ := bot.NewBotAPI(conf.GetConfig().Tgbot.Token)
		update, err := tgbot.HandleUpdate(c.Request)
		if err != nil {
			log.Error(err)
			return
		}
		updateStr, _ := json.Marshal(update)
		log.Info("get update: ", string(updateStr))
		if update.Message != nil {
			handleTgbotMessage(c, update.Message)
		}
	})
}
func handleTgbotMessage(c context.Context, m *bot.Message) {
	log := conf.GetLog(c)
	user := m.From
	chat := m.Chat
	text := m.Text
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

func checkUserChatAvailable(f func(ctx context.Context, params []string)) func(ctx context.Context, params []string) {
	return func(ctx context.Context, params []string) {
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
		msg, _ := template.TemplateGetString(template.TEMPLATE_TOO_OFTEN, nil)
		if !db.CheckUserAvailable(ctx, user) {
			sender.TgSendData(user, msg)
		}
		if !db.CheckChatAvailable(ctx, chat) {
			sender.TgSendData(user, msg)
		}
		f(ctx, params)
	}
}

func aboutCommand(ctx context.Context, params []string) {
	log := conf.GetLog(ctx)
	chat, ok := ctx.Value("tg_chat").(int64)
	if !ok {
		log.Error("cannot get telegram chat id from context")
		return
	}
	msg, _ := template.TemplateGetString(template.TEMPLATE_ABOUTME, nil)
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

	infoMsg, err := template.TemplateGetString(template.TEMPLATE_REALTIME, nil)
	if err != nil {
		log.Error("error: ", err)
		return
	}
	if len(params) != 2 {
		sender.TgSendData(chat, infoMsg)
	}

	// crypto
	item, err := db.GetCryptoItemById(ctx, params[1])
	if err != nil {
		log.Error("error: ", err)
	}
	if item != nil {
		items := []*model.GoodsItem{item}
		markets, err := fetcher.GeckoGetUsdValue(items)
		if err != nil {
			log.Info("error: ", err)
			return
		}
		output := model.Output{
			Datetime: time.Now().Format("2006-01-02 15:04:05"),
			Items:    markets,
		}
		msg, err := template.TemplateGetString(template.TEMPLATE_CRYPTO, output)
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
