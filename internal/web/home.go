package web

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

var (
	homeTemplate = template.Must(template.ParseFiles("templates/comps/header.html", "templates/comps/footer.html", "templates/comps/base.html", "templates/home.html"))
)

func RouteHome(engine *echo.Echo) {
	engine.GET("/", getHome)
}

func getHome(context echo.Context) error {
	return homeTemplate.ExecuteTemplate(context.Response(), "home.html", nil)
}
