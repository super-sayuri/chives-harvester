package router

import (
	"github.com/gin-gonic/gin"
	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sayuri_crypto_bot/conf"
)

const API_COMMAND = "api_command"

func commandRouter(r *gin.RouterGroup) {
	r.GET(routerMap[API_COMMAND], func(c *gin.Context) {
		tgbot, err := bot.NewBotAPI(conf.GetConfig().Tgbot.Token)
		if err != nil {
			ResponseError(c, 500, err)
			return
		}
		commands, err := tgbot.GetMyCommands()
		if err != nil {
			ResponseError(c, 500, err)
			return
		}
		NormalResponse(c, commands)
	})
}
