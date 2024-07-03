package main

import (
	"context"
	"go-api-test/pkg/config"
	"go-api-test/pkg/postgres"
	"log"
)

func main() {
	configPath, err := config.ParseFlags()

	if err != nil {
		log.Fatal(err)
		return
	}

	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	repositories, err := postgres.NewRepositories(config)
	if err != nil {
		log.Fatal(err)
		return
	}

	repositories.Migration(context.Background())
}
