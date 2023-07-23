package config

import (
	"filter/service/baidu"
	"github.com/spf13/viper"
	"log"
)

type config struct {
	UseBaiDu bool `mapstructure:"use_baidu"`
}

var Cfg config

func ReadConfig() {
	viper.SetConfigName("config.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
	if Cfg.UseBaiDu {
		baidu.ReadConfig()
	} else {

	}
}
