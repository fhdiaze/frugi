package cmd

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

type Command struct {
	Cmd string `form:"cmd"`
}

var (
	cmdListTemplate = template.Must(template.ParseFiles("templates/comps/frame.html", "templates/cmd/list.html"))
)

func Route(group *echo.Group) {
	group.POST("/cmd.run", handleRun)
}

func handleRun(context echo.Context) error {
	cmd := new(Command)
	context.Bind(cmd)

	return cmdListTemplate.Execute(context.Response(), cmd)
}
