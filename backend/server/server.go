package server

import "app/backend/config"

func Run() {
	port := config.GetEnv("PORT", "3001")
	router := InitializeRouter()
	err := router.Run(port)
	if err != nil {
		panic(err)
	}
}
