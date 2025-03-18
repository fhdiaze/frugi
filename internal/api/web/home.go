package web

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

var (
	homeTemplate = template.Must(template.ParseFiles("web/templates/comps/header.html", "web/templates/comps/footer.html", "web/templates/comps/base.html", "web/templates/home.html"))
)

func RouteHome(engine *echo.Echo) {
	engine.GET("/", getHome)
}

func getHome(context echo.Context) error {
	return homeTemplate.ExecuteTemplate(context.Response(), "home.html", nil)
}
