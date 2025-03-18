package web

import (
	"github.com/fhdiaze/frugi/static"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	engine := echo.New()

	Route(engine)
	AddStatic(engine)
	AddLogger(engine)
	AddCors(engine)

	engine.Logger.Fatal(engine.Start(":80"))
}

func Route(engine *echo.Echo) {
	group := engine.Group("/web")

	RouteHome(engine)
	RouteCmd(group)
	RoutePrice(group)
	RouteMoney(group)
}

func AddStatic(engine *echo.Echo) {
	engine.StaticFS("static", static.AssetsFS)
}

func AddLogger(engine *echo.Echo) {
	engine.Use(middleware.Logger())
}

func AddCors(engine *echo.Echo) {
	engine.Use(middleware.CORS())
}
