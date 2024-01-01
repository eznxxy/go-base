package main

import (
	"fmt"

	application "github.com/eznxxy/go-base"
	"github.com/eznxxy/go-base/config"
	"github.com/eznxxy/go-base/docs"
)

//	@title			Echo Base App
//	@version		1.0
//	@description	This is a base-code (ready to go).

//	@contact.name	NIX Solutions
//	@contact.url	https://www.nixsolutions.com/
//	@contact.email	ask@nixsolutions.com

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

// @BasePath	/
func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	application.Start(cfg)
}
