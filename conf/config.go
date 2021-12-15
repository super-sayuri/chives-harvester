package conf

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var _conf *Config

type Config struct {
	Common   *CommonConfig   `yml:"common"`
	Template *TemplateConfig `yml:"template"`
	Tgbot    *TgbotConfig    `yml:"tgbot"`
	Cron     *CronConfig     `yml:"cron"`
	Log      *LogConfig      `yml:"log"`
	Redis    *RedisConfig    `yml:"redis"`
	Service  *ServiceConfig  `yml:"service"`
}

type TemplateConfig struct {
	BasePath string `yml:"bathpath"`
	Crypto   string `yml:"crypto"`
}

type CommonConfig struct {
	Lang       string `yml:"lang"`
	EncKeyPath string `yml:"encryption_key_path"`
	ConfigType string `yml:"config_type"`
}

type TgbotConfig struct {
	Token      string `yml:"token"`
	Owner      string `yml:"owner"`
	CallingGap int    `yml:"call_gap"`
}

type CronConfig struct {
	Crypto string `yml:"crypto"`
}

type RedisConfig struct {
	Nodes    string `yml:"nodes"`
	Username string `yml:"username"`
	Password string `yml:"password"`
}

type ServiceConfig struct {
	Port    string `yml:"port"`
	GinMode string `yml:"gin_mode"`
}

func InitConfig(path, keyPath string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	_conf = &Config{}
	err = yaml.Unmarshal(file, _conf)
	if err != nil {
		return err
	}
	if _conf.Common.ConfigType == "file" {
		return configFromFile(keyPath)
	} else {
		return configFromFile(keyPath)
	}
}

func GetConfig() *Config {
	return _conf
}

func configFromFile(keyPath string) error {

	keys := make(map[string]string, 0)
	file, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &keys)
	if err != nil {
		return err
	}
	setKeyValues(_conf, keys)
	err = logInit(_conf.Log)
	if err != nil {
		return nil
	}
	return nil

}

func setKeyValues(conf *Config, keys map[string]string) {
	newStr, ok := keys[conf.Tgbot.Token]
	if ok {
		conf.Tgbot.Token = newStr
	}
	newStr, ok = keys[conf.Tgbot.Owner]
	if ok {
		conf.Tgbot.Owner = newStr
	}
	newStr, ok = keys[conf.Redis.Password]
	if ok {
		conf.Redis.Password = newStr
	}
}
