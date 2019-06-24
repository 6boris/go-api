package main

// title Darksteel Intelligence Computing Foundation Golang Example API
// version 1.1
// description A Golang Example Server.
// @termsOfService https://kyle.link/

// @contact.name API Support
// @contact.url https://github.com/dicf
// @contact.email kylesliu@outlook.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// host localhost:9000
// BasePath

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

import (
	"fmt"
	"github.com/dicf/go-api/routes"
	"github.com/dicf/go-api/service"
	"log"
	"net/http"
	"time"

	cfg "github.com/dicf/go-api/pkg/configure"
)

func init() {
	cfg.Setup()
	service.Setup()
}

func main() {
	router := routes.Setup()

	server := &http.Server{
		Addr:           cfg.Cfg.Server.Host + ":" + cfg.Cfg.Server.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println(cfg.Cfg)
	if err := server.ListenAndServe(); err != nil {
		log.Panicln(err.Error())
	}
}
