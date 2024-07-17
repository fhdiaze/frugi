package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Command struct {
	Cmd string `form:"cmd"`
}

func RouteCmd(engine *echo.Echo) {
	engine.POST("/cmd/run", run)
}

func run(context echo.Context) error {
	cmd := new(Command)
	context.Bind(cmd)

	return context.Render(http.StatusOK, "run", cmd)
}
