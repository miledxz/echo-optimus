package main

import (
	"golangechos/app"
	"golangechos/config"
	"golangechos/handlers"
	"golangechos/logger"
	"golangechos/services"
	"golangechos/stores"
	"log"
	"os"

	database "golangechos/db"

	"go.uber.org/zap"
)

func main() {
	_, development := config.GetConfig()
	err := logger.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(development)
	if err != nil {
		logger.Fatal("Database Connection Failure", zap.Error(err))
	}
	defer db.Close()

	s := stores.New(db)
	se := services.New(s)
	h := handlers.New(se)

	e := app.Echo()

	handlers.SetDefault(e)
	handlers.SetRouter(e, h)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5008"
	}

	logger.Fatal("Sever failure", zap.Error(e.Start(":"+PORT)))
}
