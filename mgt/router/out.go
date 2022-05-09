package router

import (
	"context"
	"os"
	"sayuri_crypto_bot/conf"
)

const (
	ApiRouter     = "router"
	CommandRouter = "command"
	PrefixRouter  = "prefix"
	TgHookRouter  = "tghook"
)

func GetUri(name string) string {
	log := conf.GetLog(context.Background())
	item, err := listOne(name)
	if err != nil {
		log.Error("Error when init router: ", err)
		os.Exit(2)
	}
	if item == nil {
		log.Error("Missing path config for: ", name)
		os.Exit(2)
	}
	return item.Value
}
