package config

import (
	"github.com/joho/godotenv"
)

var BaseURL string = "https://open.feishu.cn/open-apis"

func Init() {
	godotenv.Load()
}
