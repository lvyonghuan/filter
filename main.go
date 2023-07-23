package main

import (
	"filter/config"
	"filter/service/baidu"
	"github.com/gin-gonic/gin"
)

// 启动服务
func main() {
	config.ReadConfig()
	r := gin.Default()
	r.GET("/", handelMessage)
	r.Run(":1919")
}

// 根据配置文件处理消息
func handelMessage(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		return
	}
	message := string(data)
	var isValid bool
	if config.Cfg.UseBaiDu {
		isValid = baidu.FilterMessage(message)
	}
	var state string
	if isValid == true {
		state = "1"
	} else {
		state = "0"
	}
	c.String(200, state)
}
