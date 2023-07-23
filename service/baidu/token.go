package baidu

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

var accessToken = accessTokenResponse{}

//获取access token

func getAccessToken() {
	url := "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=" + baiduCfg.APIKey + "&client_secret=" + baiduCfg.SecretKey
	payload := strings.NewReader(``)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		log.Fatalf("%v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if accessToken.AccessToken == "" {
		log.Fatalf("获取accesstoken失败,%v", string(body))
	}
	go clock()
}

// 计时器，access token过期之前一个小时自动重申请access token
func clock() {
	timer := time.After(time.Duration(accessToken.ExpiresIn-3600) * time.Second)
	<-timer
	getAccessToken()
}
