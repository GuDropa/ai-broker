package main

import (
	"net/http"

	// "resource/apps/whatsapp"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/getChats", func(c *gin.Context) {
		// ch := whatsapp.GetChats()
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
