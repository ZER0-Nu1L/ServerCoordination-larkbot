package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"servercoordination/config"
	"servercoordination/util"

	"github.com/gin-gonic/gin"
)

type ReceiceEventBody struct {
	Schema string `json:"schema"`
	Header Header `json:"header"`
	Event  Event  `json:"event"`
}
type Header struct {
	EventID    string `json:"event_id"`
	EventType  string `json:"event_type"`
	CreateTime string `json:"create_time"`
	Token      string `json:"token"`
	AppID      string `json:"app_id"`
	TenantKey  string `json:"tenant_key"`
}
type SenderID struct {
	UnionID string `json:"union_id"`
	UserID  string `json:"user_id"`
	OpenID  string `json:"open_id"`
}
type Sender struct {
	SenderID   SenderID `json:"sender_id"`
	SenderType string   `json:"sender_type"`
	TenantKey  string   `json:"tenant_key"`
}
type ID struct {
	UnionID string `json:"union_id"`
	UserID  string `json:"user_id"`
	OpenID  string `json:"open_id"`
}
type Mentions struct {
	Key       string `json:"key"`
	ID        ID     `json:"id"`
	Name      string `json:"name"`
	TenantKey string `json:"tenant_key"`
}
type Message struct {
	MessageID   string     `json:"message_id"`
	RootID      string     `json:"root_id"`
	ParentID    string     `json:"parent_id"`
	CreateTime  string     `json:"create_time"`
	ChatID      string     `json:"chat_id"`
	ChatType    string     `json:"chat_type"`
	MessageType string     `json:"message_type"`
	Content     string     `json:"content"`
	Mentions    []Mentions `json:"mentions"`
}
type Event struct {
	Sender  Sender  `json:"sender"`
	Message Message `json:"message"`
}

func ReceiceEventHandler(c *gin.Context) {
	msgBody := ReceiceEventBody{}
	c.BindJSON(&msgBody)
	fmt.Println(msgBody)
	util.SendRespToMe(fmt.Sprintf("respond: %v\n", msgBody))

	// TODO: 应用
	GetTenantAccessToken(os.Getenv("APP_ID"), os.Getenv("APP_SECRET"))

}

type TenantAccessTokenBody struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

func GetTenantAccessToken(addID string, appSecret string) {
	url := config.BaseURL + "/auth/v3/tenant_access_token/internal"
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
	fmt.Println(string(body))
}
