package router

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/db"
	"time"
)

var routerMap map[string]string

const (
	API_PREFIX = "api_prefix"
	API_PING   = "api_ping"
	API_META   = "api_meta"
)

func InitRouter(g *gin.Engine) error {
	err := initRouterMap()
	if err != nil {
		return err
	}
	err = initCommandFuncMap()
	if err != nil {
		return err
	}
	g.Use(RequestLogger)
	r := g.Group(routerMap[API_PREFIX])

	webhookRouter(r)
	commandRouter(r)

	r.GET(routerMap[API_PING], func(c *gin.Context) {
		NormalResponse(c, "pong")
	})
	r.GET(routerMap[API_META], getRouterMap)
	return nil
}

func getRouterMap(c *gin.Context) {
	NormalResponse(c, routerMap)
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

func initRouterMap() (err error) {
	routerMap, err = db.GetApiRouter()
	if err != nil {
		return err
	}

	// validate all router
	routerKeys := []string{API_PREFIX, API_PING, API_META, API_COMMAND, API_WEBHOOK}
	missings := make([]string, 0)
	for _, routerKey := range routerKeys {
		if _, ok := routerMap[routerKey]; !ok {
			missings = append(missings, routerKey)
		}
	}
	if len(missings) != 0 {
		return errors.New(fmt.Sprintf("missing router parameters in db: %v", missings))
	}

	return nil
}
