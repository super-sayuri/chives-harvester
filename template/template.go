package template

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"sayuri_crypto_bot/conf"
	"strings"
	"text/template"
)

type Key string

const (
	Crypto   = Key("crypto")
	Aboutme  = Key("aboutme")
	Realtime = Key("realtime")
	TooOften = Key("too_often")
	Tarot    = Key("tarot")
)

var _templateMap map[Key]*template.Template

func Init(config *conf.Config) error {
	_templateMap = make(map[Key]*template.Template, 0)
	tmplToParse := []Key{Crypto, Aboutme, Realtime, TooOften, Tarot}
	for _, tmplKey := range tmplToParse {
		if err := initSingleTemplate(tmplKey); err != nil {
			return err
		}
	}
	return nil
}

func initSingleTemplate(key Key) error {
	config := conf.GetConfig()
	path := fmt.Sprintf("%s/%s_%s.tpl", config.Template.BasePath, string(key), strings.ToLower(config.Common.Lang))
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	tmpl, err := template.New(string(key)).Parse(string(data))
	if err != nil {
		return nil
	}
	_templateMap[key] = tmpl
	return nil
}

func GetString(key Key, params interface{}) (string, error) {
	tmpl, ok := _templateMap[key]
	if !ok {
		return "", errors.New("cannot find template by key")
	}
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, params)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
