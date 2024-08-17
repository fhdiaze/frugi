package web

import (
	"github.com/fhdiaze/frugi/internal/cmd"
	"github.com/fhdiaze/frugi/internal/money"
	"github.com/fhdiaze/frugi/internal/price"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(engine *echo.Echo) {
	group := engine.Group("/web")

	RouteHome(engine)
	cmd.Route(group)
	price.Route(group)
	money.Route(group)
}

func AddStatic(engine *echo.Echo) {
	engine.Static("/assets", "assets")
}

func AddLogger(engine *echo.Echo) {
	engine.Use(middleware.Logger())
}
