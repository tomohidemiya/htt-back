package server

import (
	health "app/backend/controllers"
	"app/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func InjectRouting(engine *gin.Engine) {
	// Anyone
	engine.GET("/ping", health.CheckHertBeat)

	// Need Authentication
	engine.Use(middlewares.AuthMiddleware())
}
