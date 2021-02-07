package server

import (
	"app/backend/controllers"
	"app/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func InjectRouting(engine *gin.Engine) {
	// Anyone
	healthCtrl := health.Controller{}
	engine.GET("/ping", healthCtrl.Index)

	// Need Authentication
	engine.Use(middlewares.AuthMiddleware())
}
