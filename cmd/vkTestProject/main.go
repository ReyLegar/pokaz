package main

import (
	"context"
	"flag"
	"log"

	"github.com/ReyLegar/vkTestProject/internal/app"
	"github.com/ReyLegar/vkTestProject/internal/config"
	"github.com/ReyLegar/vkTestProject/internal/handler"
	"github.com/ReyLegar/vkTestProject/internal/repository"
	"github.com/ReyLegar/vkTestProject/internal/service"
)

func main() {

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.NewConfig()

	a, err := app.NewApp(ctx, cfg)

	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.ConnectDB(cfg)

	if err != nil {
		log.Fatal("Error!")
	}

	repos := repository.NewRepository(db)
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)
	handlers.InitRoutes()

	err = a.Run(handlers)

	if err != nil {
		log.Fatal(err)
	}

}
