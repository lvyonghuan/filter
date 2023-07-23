package baidu

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type data struct {
	LogID          int    `json:"log_id"`
	Conclusion     string `json:"conclusion"`
	ConclusionType int    `json:"conclusionType"`
	ErrorMsg       string `json:"error_msg"`
}

//进行过滤

func FilterMessage(message string) bool {
	reqUrl := "https://aip.baidubce.com/rest/2.0/solution/v1/text_censor/v2/user_defined?access_token=" + accessToken.AccessToken

	params := url.Values{}
	params.Set("text", message)

	resp, err := http.PostForm(reqUrl, params)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()
	if resp.Body == nil {
		log.Println("err")
		return false
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	var response = data{}
	err = json.Unmarshal(body, &response)
	if response.ErrorMsg != "" {
		log.Println("response err:", string(body))
		return false
	}
	if response.ConclusionType == 1 {
		return true
	} else { //最高等级过滤，其实可以细化，但是懒得了
		return false
	}
}
