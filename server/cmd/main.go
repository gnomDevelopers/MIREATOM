package main

import (
	"fmt"
	"server/internal/handler"
	logger "server/internal/log"
	"server/internal/repository/postgres"
)

func main() {
	log := logger.InitLogger()
	db, err := postgres.NewDatabase()
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("could not initialize database connection: %s", err))
	}

	handlers := handler.NewHandler(db, log)

	app := handlers.Router()
	app.Listen(":8080")
}
