package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Encrypt Key模式
var ifEncrypt = false
var encryptKey string = "sdnabcdefg" // TODO: 更改并写成配置文件
func Decrypt(encrypt string, key string) (string, error) {
	buf, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", fmt.Errorf("base64StdEncode Error[%v]", err)
	}
	if len(buf) < aes.BlockSize {
		return "", errors.New("cipher  too short")
	}
	keyBs := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(keyBs[:sha256.Size])
	if err != nil {
		return "", fmt.Errorf("AESNewCipher Error[%v]", err)
	}
	iv := buf[:aes.BlockSize]
	buf = buf[aes.BlockSize:]
	// CBC mode always works in whole blocks.
	if len(buf)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(buf, buf)
	n := strings.Index(string(buf), "{")
	if n == -1 {
		n = 0
	}
	m := strings.LastIndex(string(buf), "}")
	if m == -1 {
		m = len(buf) - 1
	}
	return string(buf[n : m+1]), nil
}

type challengeBodyInEncrypt struct {
	// NOTE: 实际上格式更复杂，文档上说得有歧义，有这个模式的需求，请根据https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM?lang=zh-CN#2eb3504a再更改
	Encrypt string `json:"encrypt"`
}

type challengeBody struct {
	Challenge string `json:"challenge"`
	Token     string `json:"token"`
	Type      string `json:"type"`
}

func challengeInEncryptHandler(c *gin.Context) {
	challBody := challengeBodyInEncrypt{}
	c.BindJSON(&challBody)
	encrypt := challBody.Encrypt

	s, err := Decrypt(encrypt, encryptKey)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"challenge": s,
	})
}
func challengeHandler(c *gin.Context) {
	challBody := challengeBody{}
	c.BindJSON(&challBody)

	c.JSON(http.StatusOK, gin.H{
		"challenge": challBody.Challenge,
	})
}

func main() {
	r := gin.Default()
	if ifEncrypt {
		r.POST("/", challengeInEncryptHandler)
	} else {
		r.POST("/", challengeHandler)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // [default]listen and serve on 0.0.0.0:8080
}
