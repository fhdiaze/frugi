package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(engine *echo.Echo) {
	group := engine.Group("/web")
	RouteHome(engine)
	RouteCmd(group)
	RoutePrice(group)
	RouteMoney(group)
}

func AddStatic(engine *echo.Echo) {
	engine.Static("/assets", "assets")
}

func AddLogger(engine *echo.Echo) {
	engine.Use(middleware.Logger())
}
