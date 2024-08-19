package main

import (
	"context"
	"github.com/RandySteven/neo-postman/apps"
	"github.com/RandySteven/neo-postman/enums"
	"github.com/RandySteven/neo-postman/middlewares"
	"github.com/RandySteven/neo-postman/pkg/config"
	"github.com/RandySteven/neo-postman/pkg/elastics"
	"github.com/RandySteven/neo-postman/pkg/postgres"
	"github.com/RandySteven/neo-postman/pkg/redis"
	"github.com/RandySteven/neo-postman/pkg/scheduler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	ctx2 := context.WithValue(ctx, enums.ActivePostgres, 1)
	repositories, err := postgres.NewRepositories(config)
	if err != nil {
		ctx2 = context.WithValue(ctx2, enums.ActivePostgres, 0)
		log.Fatal(err)
	}

	ctx3 := context.WithValue(ctx2, enums.ActiveRedis, 1)
	caches, err := redis.NewRedis(config)
	if err != nil {
		ctx3 = context.WithValue(ctx3, enums.ActiveRedis, 0)
		log.Fatal(err)
	}

	err = caches.Ping(ctx3)
	if err != nil {
		ctx3 = context.WithValue(ctx3, enums.ActiveRedis, 0)
		log.Fatal(err)
	}

	ctx4 := context.WithValue(ctx3, enums.ActiveElastic, 1)
	documentaries, err := elastics.NewESClient(config)
	if err != nil {
		ctx4 = context.WithValue(ctx4, enums.ActiveElastic, 0)
		log.Fatal(err)
	}
	if err = documentaries.Ping(ctx3); err != nil {
		ctx4 = context.WithValue(ctx4, enums.ActiveElastic, 0)
		log.Fatal(err)
	}

	schedulerAct := scheduler.NewScheduler(*repositories, *caches)
	err = schedulerAct.RunAllJob(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ctx = ctx4

	handlers := apps.NewHandlers(repositories, caches, documentaries)
	r := mux.NewRouter()
	r = apps.RegisterMiddleware(r)
	r.Use(middlewares.ContextMiddleware(ctx))
	handlers.InitRouter(r)

	go config.Run(r)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = caches.ClearCache(ctx); err != nil {
		log.Fatal(err)
		return
	}
}
