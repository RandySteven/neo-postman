package main

import (
	"context"
	"github.com/RandySteven/neo-postman/apps"
	"github.com/RandySteven/neo-postman/pkg/config"
	"github.com/RandySteven/neo-postman/pkg/postgres"
	"github.com/gorilla/mux"
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
	r = apps.RegisterMiddleware(r)
	log.Println("UDAH KE RUN WA")
	handlers.InitRouter(r)
	config.Run(r)
}
