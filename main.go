package main

import (
	"app/config"
	"app/db"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	db.Init()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}