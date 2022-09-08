package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"servercoordination/config"
)

type SendBody struct {
	MsgType string  `json:"msg_type"`
	Content Content `json:"content"`
}
type Content struct {
	Text string `json:"text"`
}

func SendRespToMe(msg string) {
	data := SendBody{}
	data.MsgType = "text"
	data.Content.Text = msg

	payload, _ := json.Marshal(data)

	client := &http.Client{}
	req, err := http.NewRequest("POST", config.LarkHook, bytes.NewReader(payload))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
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
