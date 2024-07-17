package main

import (
	"github.com/fhdiaze/frugi/api"
	"github.com/labstack/echo/v4"
)

func main() {
	engine := echo.New()

	api.Route(engine)
	api.AddStatic(engine)
	api.AddRender(engine)
	api.AddLogger(engine)

	engine.Logger.Fatal(engine.Start(":8080"))
}
