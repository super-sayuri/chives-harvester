package util

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"sayuri_crypto_bot/conf"
	"strings"
	"text/template"
)

type TemplateKey string

const (
	TEMPLATE_CRYPTO    = TemplateKey("crypto")
	TEMPLATE_ABOUTME   = TemplateKey("aboutme")
	TEMPLATE_REALTIME  = TemplateKey("realtime")
	TEMPLATE_TOO_OFTEN = TemplateKey("too_often")
)

var _templateMap map[TemplateKey]*template.Template

func templateInit(config *conf.Config) error {
	_templateMap = make(map[TemplateKey]*template.Template, 0)
	tmplToParse := []TemplateKey{TEMPLATE_CRYPTO, TEMPLATE_ABOUTME, TEMPLATE_REALTIME, TEMPLATE_TOO_OFTEN}
	for _, tmplKey := range tmplToParse {
		if err := initSingleTemplate(tmplKey); err != nil {
			return err
		}
	}
	return nil
}

func initSingleTemplate(key TemplateKey) error {
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
