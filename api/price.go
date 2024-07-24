package api

import (
	"html/template"
	"net/http"

	"github.com/fhdiaze/frugi/module/price"
	"github.com/labstack/echo/v4"
)

var (
	scaleGetTemplate = template.Must(template.ParseFiles("templates/comps/frame.html", "templates/price/scale.html"))
	scaleOutTemplate = template.Must(template.ParseFiles("templates/comps/frame.html", "templates/comps/result.html", "templates/price/scale_out.html"))
)

func RoutePrice(engine *echo.Echo) {
	engine.GET("/price.scale.get", handleGetScale)
	engine.POST("/price.scale.run", handleRunScale)
}

func handleGetScale(context echo.Context) error {
	return scaleGetTemplate.ExecuteTemplate(context.Response(), "scale.html", nil)
}

func handleRunScale(context echo.Context) error {
	var cmd price.RunScaleCmd
	err := context.Bind(&cmd)

	if err != nil {
		return context.String(http.StatusBadRequest, "bad request")
	}

	u, err := price.HandleRunScale(&cmd)

	if err != nil {
		return err
	}

	return scaleOutTemplate.ExecuteTemplate(context.Response(), "scale_out.html", u)
}
