package baidu

import (
	"github.com/spf13/viper"
	"log"
)

type baiduConfig struct {
	APIKey    string `mapstructure:"api_key"`
	SecretKey string `mapstructure:"secret_key"`
}

var baiduCfg baiduConfig

func ReadConfig() {
	viper.SetConfigName("baidu_config.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&baiduCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
	getAccessToken()
}
