package main

import (
	"fmt"

	application "github.com/eznxxy/go-base"
	"github.com/eznxxy/go-base/config"
	"github.com/eznxxy/go-base/docs"
)

// @title Echo Demo App
// @version 1.0
// @description This is a demo version of Echo app.

// @contact.name NIX Solutions
// @contact.url https://www.nixsolutions.com/
// @contact.email ask@nixsolutions.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /
func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.ExposePort)

	application.Start(cfg)
}
