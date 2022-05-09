package command

import (
	"github.com/gin-gonic/gin"
	"sayuri_crypto_bot/mgt/router"
	"sayuri_crypto_bot/util"
)

func ImplementRouter(g *gin.RouterGroup) {
	baseUri := router.GetUri(router.CommandRouter)
	g.GET(baseUri, func(c *gin.Context) {
		items, err := read(c)
		if err != nil {
			util.ResponseError(c, 500, err)
		} else {
			util.NormalResponse(c, items)
		}
	})
	g.POST(baseUri+"/refresh", func(c *gin.Context) {
		items, err := updateCache(c)
		if err != nil {
			util.ResponseError(c, 500, err)
		} else {
			util.NormalResponse(c, items)
		}
	})
}
