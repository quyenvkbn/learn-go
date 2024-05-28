package main

import (
	"learn-go/configs"
	"learn-go/seeds"
	"log"
)

func main() {
	if seeds.Seeds() {
		return
	}

	router, err := InitializeController()
	server := router.SetupRoutes()
	if err != nil {
		log.Fatal(err)
	}

	configs.LoadConfigJson(".")
	configApplication := configs.GetConfigByKey("application")
	log.Println("Starting server on :" + configApplication["port"].(string))
	if err := server.Run(":" + configApplication["port"].(string)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
