package server

import (
	"fmt"

	validatorV10 "github.com/go-playground/validator/v10"
	settings "github.com/hjcalderon10/bunny-backend/setting"
	"github.com/hjcalderon10/bunny-backend/util/validator"
	"github.com/labstack/echo"
)

func createServer() *echo.Echo {
	server := echo.New()

	server.Debug = settings.Commons.AppEnv == "development"
	server.HideBanner = true
	server.Validator = &validator.ApiValidator{Validator: validatorV10.New()}

	SetupMiddlewares(server)
	SetupRoutes(server)

	return server

}

func Start() {
	s := createServer()
	host, port := settings.Commons.Host, settings.Commons.Port

	s.Logger.Fatal(s.Start(fmt.Sprintf("%s:%s", host, port)))
}
