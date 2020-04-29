package common

import (
	"github.com/kyour-cn/gourd/application"
	"github.com/kyour-cn/gourd/utils/toml"
	"log"
)

var configFile string

//读取配置信息
func ReadConfig(file string) (config *application.Config) {

	if configFile == "" {
		configFile = file
	}

	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}

	return
}
