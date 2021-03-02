package wechat

import (
	"log"
	"strings"
	"wejh-go/config"
	"wejh-go/exception"
)

type wechatConfig struct {
	Driver    string
	AppId     string
	AppSecret string
}

func getConfigs() wechatConfig {

	wc := wechatConfig{}
	if !config.Config.IsSet("wechat.appid") {
		log.Fatal(exception.ConfigError)
	}
	if !config.Config.IsSet("wechat.appsecret") {
		log.Fatal(exception.ConfigError)
	}
	wc.AppId = config.Config.GetString("wechat.appid")
	wc.AppSecret = config.Config.GetString("wechat.appsecret")

	wc.Driver = string(Memory)
	if config.Config.IsSet("wechat.driver") {
		wc.Driver = strings.ToLower(config.Config.GetString("wechat.driver"))
	}
	return wc
}
