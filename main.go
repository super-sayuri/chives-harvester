package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/db"
	"sayuri_crypto_bot/fetcher/crypto"
	"sayuri_crypto_bot/fortune"
	"sayuri_crypto_bot/job"
	"sayuri_crypto_bot/router"
	"sayuri_crypto_bot/sender"
	"sayuri_crypto_bot/template"
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

	err = template.Init(conf.GetConfig())
	if err != nil {
		log.Fatal("error when init util: ", err)
	}

	err = db.Init(conf.GetConfig())
	if err != nil {
		log.Fatal("error when init db: ", err)
	}

	err = fortune.Init(conf.GetConfig())
	if err != nil {
		log.Fatal("error when init fortune: ", err)
	}

	err = crypto.InitCryptoFetcher(conf.GetConfig().Crypto)
	if err != nil {
		log.Fatal("error when init crypto fetcher: ", err)
	}

	job.CronInit()
	sender.TgStartMessage(conf.GetConfig().Tgbot.Owner)

	router.HandleCommands()
}
