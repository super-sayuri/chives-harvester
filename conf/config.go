package conf

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var _conf *Config

type Config struct {
	Common   *CommonConfig   `yaml:"common"`
	Template *TemplateConfig `yaml:"template"`
	Tgbot    *TgbotConfig    `yaml:"tgbot"`
	Cron     *CronConfig     `yaml:"cron"`
	Log      *LogConfig      `yaml:"log"`
	Redis    *RedisConfig    `yaml:"redis"`
	Service  *ServiceConfig  `yaml:"service"`
}

type TemplateConfig struct {
	BasePath string `yaml:"basepath"`
	Crypto   string `yaml:"crypto"`
}

type CommonConfig struct {
	Lang       string `yaml:"lang"`
	EncKeyPath string `yaml:"encryption_key_path"`
	ConfigType string `yaml:"config_type"`
	Name       string `yaml:"name"`
}

type TgbotConfig struct {
	Token      string `yaml:"token"`
	Owner      string `yaml:"owner"`
	CallingGap int    `yaml:"call_gap"`
}

type CronConfig struct {
	Crypto string `yaml:"crypto"`
}

type RedisConfig struct {
	Nodes    string `yaml:"nodes"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ServiceConfig struct {
	Port    string `yaml:"port"`
	GinMode string `yaml:"gin_mode"`
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
