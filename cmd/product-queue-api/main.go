package main

import (
	"log"

	"github.com/ikiwq/go-inflation/internal/product-queue-api/api"
	"github.com/ikiwq/go-inflation/internal/product-queue-api/config"
)

func main() {
	config, err := config.NewProductQueueApiConfiguration("config.yml")
	if err != nil {
		log.Fatal("error while opening config.yaml:", err)
	}

	api := api.NewApi(*config)
	api.Start()

	defer api.Exit()
}
