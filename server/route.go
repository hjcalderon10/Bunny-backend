package server

import (
	"net/http"
	"time"

	models "github.com/hjcalderon10/bunny-backend/model"
	settings "github.com/hjcalderon10/bunny-backend/setting"
	"github.com/labstack/echo"
)

const BASE_PATH = "api"

func SetupRoutes(server *echo.Echo) {
	baseRout := server.Group(BASE_PATH)

	baseRout.GET("/health-check", Get)
}

func Get(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &models.HealthCheck{
		Name:    settings.Commons.ProjectName,
		Version: settings.Commons.ProjectVersion,
		Date:    time.Now().UTC(),
	})
}
