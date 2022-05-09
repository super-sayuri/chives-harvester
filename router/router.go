package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/mgt/command"
	"sayuri_crypto_bot/mgt/router"
	"time"
)

//var routerMap map[string]string

func InitRouter(g *gin.Engine) (err error) {
	if err != nil {
		return err
	}
	err = initCommandFuncMap()
	if err != nil {
		return err
	}
	g.Use(RequestLogger)
	r := g.Group(router.GetUri(router.PrefixRouter))

	webhookRouter(r)

	command.ImplementRouter(r)
	router.ImplementRouter(r)
	return nil
}

func RequestLogger(c *gin.Context) {
	reqId := uuid.NewString()
	c.Set(conf.LOG_KEY_REQUEST_ID, reqId)
	l := conf.GetLog(c)
	l.Info("Req #", reqId, " starts.")
	// url info
	l.Info("Calling ", c.Request.Method, " ", c.Request.RequestURI)
	// user info
	l.Info("User ip:", c.ClientIP(), ", agent:", c.Request.Header.Get("User-Agent"))
	c.Next()
	l.Info("Result: code ", c.Writer.Status(), ", size: ", c.Writer.Size())
	l.Info("Req #", reqId, " ends at ", time.Now())
}
