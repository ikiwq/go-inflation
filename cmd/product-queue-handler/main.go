package main

import (
	"log"

	"github.com/ikiwq/go-inflation/internal/product-queue-handler/config"
	"github.com/ikiwq/go-inflation/internal/product-queue-handler/handler"
)

func main() {
	config, err := config.NewProductQueueHandlerConfig("config.yml")
	if err != nil {
		log.Fatal("error while opening config.yaml:", err)
	}

	handler := handler.NewHandler(config)
	handler.Start()

	defer handler.Exit()
}
