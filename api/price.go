package api

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

var (
	priceScaleGetTemplate = template.Must(template.ParseFiles("templates/comps/frame.html", "templates/price/scale.html"))
)

func RoutePrice(engine *echo.Echo) {
	engine.GET("/price.scale.get", handleScalePriceGet)
}

func handleScalePriceGet(context echo.Context) error {
	return priceScaleGetTemplate.ExecuteTemplate(context.Response(), "scale.html", nil)
}
