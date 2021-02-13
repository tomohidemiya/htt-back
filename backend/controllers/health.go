package health

import (
	"github.com/gin-gonic/gin"
)

type Pong struct {
	Message string
}

func CheckHertBeat(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
