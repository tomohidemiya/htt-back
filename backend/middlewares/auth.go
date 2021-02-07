package middlewares

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}