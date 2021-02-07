package main

import (
	"app/backend/config"
	"app/backend/db"
	"app/backend/server"
)

func main() {
	config.Init()
	db.Init()
	server.Run()
}