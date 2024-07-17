package api

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, context echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Route(engine *echo.Echo) {
	RouteHome(engine)
	RouteCmd(engine)
}

func AddStatic(engine *echo.Echo) {
	engine.Static("/assets", "assets")
}

func AddRender(engine *echo.Echo) {
	t := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	engine.Renderer = t
}

func AddLogger(engine *echo.Echo) {
	engine.Use(middleware.Logger())
}
