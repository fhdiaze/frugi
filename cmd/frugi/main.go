package main

import (
	"github.com/fhdiaze/frugi/internal/web"
	"github.com/labstack/echo/v4"
)

func main() {
	engine := echo.New()

	web.Route(engine)
	web.AddStatic(engine)
	web.AddLogger(engine)

	engine.Logger.Fatal(engine.Start(":8080"))
}
