package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var BaseURL string = "https://open.feishu.cn/open-apis"

type TenantAccessTokenBody struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type TenantAccessTokenReply struct {
	Code              int    `json:"code"`
	Expire            int    `json:"expire"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
}

func GetTenantAccessToken(addID string, appSecret string) (token string) {
	url := BaseURL + "/auth/v3/tenant_access_token/internal"
	method := "POST"

	data := TenantAccessTokenBody{
		AppID:     addID,
		AppSecret: appSecret,
	}
	payload, _ := json.Marshal(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("Cookie", "Cookie_2=value")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	respond := TenantAccessTokenReply{}

	// marshal
	err = json.Unmarshal(body, &respond)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Marshal Body: %v, error: %v", body, err)
	}
	return respond.TenantAccessToken // TODO: err
}

var AccessToken string

func InitConf() {
	godotenv.Load()

	AppID := os.Getenv("APP_ID")
	AppSecret := os.Getenv("APP_SECRET")
	AccessToken = GetTenantAccessToken(AppID, AppSecret)
}
