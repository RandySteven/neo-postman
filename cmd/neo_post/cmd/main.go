package main

import (
	"context"
	"github.com/RandySteven/neo-postman/apps"
	"github.com/RandySteven/neo-postman/pkg/config"
	"github.com/RandySteven/neo-postman/pkg/postgres"
	"github.com/RandySteven/neo-postman/pkg/redis"
	"github.com/RandySteven/neo-postman/pkg/scheduler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load("./files/env/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
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

	caches, err := redis.NewRedis(config)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = caches.Ping(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	scheduler := scheduler.NewScheduler(*repositories)
	err = scheduler.RunAllJob(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	handlers := apps.NewHandlers(repositories, caches)
	r := mux.NewRouter()
	r = apps.RegisterMiddleware(r)

	log.Println("UDAH KE RUN WA")
	handlers.InitRouter(r)
	config.Run(r)
}
