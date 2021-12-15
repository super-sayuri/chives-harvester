package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/db"
	"sayuri_crypto_bot/job"
	"sayuri_crypto_bot/router"
	"sayuri_crypto_bot/sender"
	"sayuri_crypto_bot/util"
)

var (
	confPath string
	keyPath  string
)

func init() {
	flag.StringVar(&confPath, "c", "", "config file path")
	flag.StringVar(&keyPath, "k", "keys.json", "config file path")
}

func main() {
	flag.Parse()
	var err error
	err = conf.InitConfig(confPath, keyPath)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(200)
	}

	err = util.Init(conf.GetConfig())
	if err != nil {
		log.Fatal("error when init util: ", err)
	}

	err = db.Init(conf.GetConfig())
	if err != nil {
		log.Fatal("error when init db: ", err)
	}

	job.CronInit()
	sender.TgStartMessage(conf.GetConfig().Tgbot.Owner)
	gin.SetMode(conf.GetConfig().Service.GinMode)
	g := gin.Default()
	g.SetTrustedProxies(nil)
	err = router.InitRouter(g)
	if err != nil {
		log.Fatal("error when init service: ", err)
	}
	g.Run(fmt.Sprintf(":%s", conf.GetConfig().Service.Port))
}
