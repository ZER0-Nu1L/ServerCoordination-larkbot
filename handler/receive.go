package handler

import (
	"fmt"

	"servercoordination/config"

	"github.com/cihub/seelog"
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
	// util.SendRespToMe(fmt.Sprintf("respond: %v\n", msgBody)) // DEBUG:

	chatID := msgBody.Event.Message.ChatID
	// chatID := "oc_020b35e61a3bf471ead260a3c586f184" // DEBUG:

	// TODO: add feature
	content := "{\"text\":\"<at user_id=\\\"dbg369f5\\\">吴昌博</at> test success\"}"
	seelog.Info(chatID, config.AccessToken, content)
	Send2Chat(config.AccessToken, chatID, content, "text")
}
