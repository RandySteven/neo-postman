package main

import (
	"context"
	"github.com/gorilla/mux"
	"go-api-test/apps"
	"go-api-test/pkg/config"
	"go-api-test/pkg/postgres"
	"log"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
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

	handlers := apps.NewHandlers(repositories)
	r := mux.NewRouter()
	log.Println("UDAH KE RUN WA")
	handlers.InitRouter(r)
	config.Run(r)
}
