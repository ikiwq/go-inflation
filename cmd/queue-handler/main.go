package main

import (
	"log"

	"github.com/ikiwq/go-inflation/internal/queue-handler/config"
	"github.com/ikiwq/go-inflation/internal/queue-handler/handler"
)

func main() {
	config, err := config.NewQueueHandlerConfig("config.yml")
	if err != nil {
		log.Fatal("error while opening config.yaml:", err)
	}

	handler := handler.NewHandler(config)
	handler.Start()

	defer handler.Exit()
}
