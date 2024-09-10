package sender

import (
	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"sayuri_crypto_bot/collections"
	"sayuri_crypto_bot/conf"
	"strconv"
	"sync"
)

var getTgBot = sync.OnceValues(func() (*bot.BotAPI, error) {
	return bot.NewBotAPI(conf.GetConfig().Tgbot.Token)
})

func TgSendData(userId int64, msg string) error {
	tgbot, err := getTgBot()
	if err != nil {
		return err
	}
	c := bot.MessageConfig{

		BaseChat: bot.BaseChat{
			ChatID: userId,
		},
		Text: msg,
	}
	_, err = tgbot.Send(c)
	if err != nil {
		return err
	}
	return nil
}

func TgStartMessage(ownerId string) error {
	id, err := strconv.Atoi(ownerId)
	if err != nil {
		return err
	}
	tgbot, err := getTgBot()
	if err != nil {
		return err
	}
	c := bot.MessageConfig{

		BaseChat: bot.BaseChat{
			ChatID: int64(id),
		},
		Text: "天灾再起! " + conf.GetConfig().Common.Name,
	}
	_, err = tgbot.Send(c)
	if err != nil {
		return err
	}
	return nil

}

type stickerName string

func (s stickerName) NeedsUpload() bool {
	return false
}

func (s stickerName) UploadData() (string, io.Reader, error) {
	return "", nil, nil
}

func (s stickerName) SendData() string {
	return string(s)
}

var _ bot.RequestFileData = (*stickerName)(nil)

func TgSendSticker(userId int64) error {
	tgbot, err := getTgBot()
	if err != nil {
		return err
	}

	// get random sticker from FXSENSHI
	stickers, err := getStickerNames("FXSENSHI")
	if err != nil {
		return err
	}

	// get random sticker
	stickerFile := collections.SliceRandomPick(stickers)

	c := bot.NewSticker(userId, stickerFile)
	_, err = tgbot.Send(c)
	return err
}

func getStickerNames(setName string) ([]stickerName, error) {
	// get it from cache
	// or else get from telegram api
	tgbot, err := getTgBot()
	if err != nil {
		return nil, err
	}

	stickerData, err := tgbot.GetStickerSet(bot.GetStickerSetConfig{Name: "FXSENSHI"})
	if err != nil {
		return nil, err
	}

	stickerNames := make([]stickerName, len(stickerData.Stickers))

	for i, sticker := range stickerData.Stickers {
		stickerNames[i] = stickerName(sticker.FileID)
	}

	return stickerNames, nil
}
