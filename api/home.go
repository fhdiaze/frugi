package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RouteHome(engine *echo.Echo) {
	engine.GET("/", getHome)
}

func getHome(context echo.Context) error {
	return context.Render(http.StatusOK, "home.html", nil)
}
