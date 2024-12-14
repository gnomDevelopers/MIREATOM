package main

import (
	"fmt"
	_ "server/docs"
	"server/internal/handler"
	logger "server/internal/log"
	"server/internal/repository/postgres"
	"server/util"
)

// @title Σigma API
// @version 1.0
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log := logger.InitLogger()
	db, err := postgres.NewDatabase()
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("could not initialize database connection: %s", err))
	}
	util.CreateTmp()
	handlers := handler.NewHandler(db, log)

	app := handlers.Router()
	app.Listen(":8080")
}
