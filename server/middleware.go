package server

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetupMiddlewares(server *echo.Echo) {
	server.Use(
		middleware.Recover(),
		middleware.RequestID(),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "[method:${method}][uri:${uri}][status:${status}][latency_human:${latency_human}][xAppID:${header:x-application-id}]\n",
			Skipper: func(e echo.Context) bool {
				return strings.Contains(e.Path(), "health-check")
			},
		}),
	)
}
