package main

import (
	"net/http"

	"servercoordination/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// if handler.IfEncrypt {
	// 	r.POST("/", handler.ChallengeInEncryptHandler)
	// } else {
	// 	r.POST("/", handler.ChallengeHandler)
	// }
	r.POST("/", handler.ReceiceEventHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:9000")
}
