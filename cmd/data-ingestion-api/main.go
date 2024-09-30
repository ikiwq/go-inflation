package main

import (
	"log"

	"github.com/ikiwq/go-inflation/internal/data-ingestion-api/api"
	"github.com/ikiwq/go-inflation/internal/data-ingestion-api/config"
)

func main() {
	config, err := config.NewDataIngestionApiConfiguration("config.yml")
	if err != nil {
		log.Fatal("error while opening config.yaml:", err)
	}

	api := api.NewApi(*config)
	api.Start()

	defer api.Exit()
}
