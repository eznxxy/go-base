package application

import (
	"log"

	"github.com/eznxxy/go-base/config"
	"github.com/eznxxy/go-base/server"
	"github.com/eznxxy/go-base/server/routes"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
