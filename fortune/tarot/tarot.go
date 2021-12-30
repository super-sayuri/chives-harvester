package tarot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sayuri_crypto_bot/conf"
	"time"
)

type TarotCard struct {
	Name string   `json:"name"`
	Des  []string `json:"des"`
}

type TarotData struct {
	Position []string     `json:"position"`
	Cards    []*TarotCard `json:"cards"`
}

type TarotDraw struct {
	Position string `json:"position"`
	Name     string `json:"name"`
	Des      string `json:"des"`
}

var _tarotData *TarotData

func Init(config *conf.Config) error {
	path := fmt.Sprintf("%s/tarot_%s.json", config.Template.BasePath, config.Common.Lang)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	_tarotData = &TarotData{}
	err = json.Unmarshal(data, _tarotData)
	if err != nil {
		return err
	}
	return nil
}

func Draw() *TarotDraw {
	rand.Seed(time.Now().Unix())
	card := rand.Intn(len(_tarotData.Cards))
	pos := rand.Intn(len(_tarotData.Position))
	return &TarotDraw{
		Position: _tarotData.Position[pos],
		Name:     _tarotData.Cards[card].Name,
		Des:      _tarotData.Cards[card].Des[pos],
	}
}
