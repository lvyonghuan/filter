package main

import (
	"filter/config"
	"filter/service/baidu"
	"log"
	"testing"
)

func TestReadConfig(t *testing.T) {
	config.readConfig()
	log.Println(config.cfg)
}

func TestFilterMessage(t *testing.T) {
	baidu.ReadConfig()
	log.Println(baidu.FilterMessage("你好"))
	log.Println(baidu.FilterMessage("我是你爹"))
}
