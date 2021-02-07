package main

import (
	"app/backend/config"
	"app/backend/db"
	"app/backend/server"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	db.Init()
	defer db.Close()

	r := gin.Default()
	server.InjectRouting(r)
	r.Run()
}