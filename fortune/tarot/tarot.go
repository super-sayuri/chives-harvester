package tarot

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sayuri_crypto_bot/conf"
	"time"
)

type Card struct {
	Name string   `json:"name"`
	Des  []string `json:"des"`
}

type Data struct {
	Position []string `json:"position"`
	Cards    []*Card  `json:"cards"`
}

type Destiny struct {
	Position string `json:"position"`
	Name     string `json:"name"`
	Des      string `json:"des"`
}

var _tarotData *Data

func Init(config *conf.Config) error {
	path := fmt.Sprintf("%s/tarot_%s.json", config.Template.BasePath, config.Common.Lang)
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	_tarotData = &Data{}
	err = json.Unmarshal(data, _tarotData)
	if err != nil {
		return err
	}
	return nil
}

func Draw() *Destiny {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	card := r.Intn(len(_tarotData.Cards))
	pos := r.Intn(len(_tarotData.Position))
	return &Destiny{
		Position: _tarotData.Position[pos],
		Name:     _tarotData.Cards[card].Name,
		Des:      _tarotData.Cards[card].Des[pos],
	}
}
