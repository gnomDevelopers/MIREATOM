package main

import (
	"fmt"
	_ "server/docs"
	"server/internal/config"
	"server/internal/handler"
	logger "server/internal/log"
	"server/internal/repository/postgres"
)

// @title Σigma API
// @version 1.0
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log := logger.InitLogger()
	log.Println(config.DBHost, config.DBPort, config.DBName, config.DBUser, config.DBPassword)
	db, err := postgres.NewDatabase()
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("could not initialize database connection: %s", err))
	}

	handlers := handler.NewHandler(db, log)

	app := handlers.Router()
	app.Listen(":8080")
}
