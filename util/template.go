package util

import (
	"bytes"
	"errors"
	"io/ioutil"
	"sayuri_crypto_bot/conf"
	"text/template"
)

type TemplateKey string

const (
	CRYPTO_TEMPLATE = TemplateKey("crypto")
)

var _templateMap map[TemplateKey]*template.Template

func templateInit(config *conf.Config) error {
	_templateMap = make(map[TemplateKey]*template.Template, 0)
	cryptoPath := config.Template.Crypto
	//tmpl, err := template.New(cryptoPath).ParseFiles(cryptoPath)
	data, err := ioutil.ReadFile(cryptoPath)
	if err != nil {
		return err
	}
	tmpl, err := template.New(string(CRYPTO_TEMPLATE)).Parse(string(data))
	if err != nil {
		return err
	}
	_templateMap[CRYPTO_TEMPLATE] = tmpl
	return nil
}

func TemplateGetString(key TemplateKey, params interface{}) (string, error) {
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
