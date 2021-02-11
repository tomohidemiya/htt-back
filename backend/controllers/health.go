package health

import (
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (pc Controller) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
