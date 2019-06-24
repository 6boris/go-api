package routes

import (
	"github.com/dicf/go-api/docs"
	"github.com/dicf/go-api/pkg/configure"
)

func SetupSwagger() {
	cfg := configure.Cfg
	docs.SwaggerInfo.Title = cfg.App.Title
	docs.SwaggerInfo.Description = cfg.App.Desc
	docs.SwaggerInfo.Version = cfg.App.Version
	docs.SwaggerInfo.Host = cfg.Server.Host + ":" + cfg.Server.Port
	//docs.SwaggerInfo.BasePath = "/api"
}
