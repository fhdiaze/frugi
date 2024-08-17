package price

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	scaleGetTemplate = template.Must(template.ParseFiles("templates/comps/frame.html", "templates/price/scale.html"))
	scaleOutTemplate = template.Must(template.ParseFiles("templates/comps/frame.html", "templates/comps/result.html", "templates/price/scale_out.html"))
)

func Route(group *echo.Group) {
	group.GET("/price.scale.get", handleGetScale)
	group.POST("/price.scale.run", handleRunScale)
}

func handleGetScale(context echo.Context) error {
	return scaleGetTemplate.ExecuteTemplate(context.Response(), "scale.html", nil)
}

func handleRunScale(context echo.Context) error {
	var cmd RunScaleCmd
	err := context.Bind(&cmd)

	if err != nil {
		return context.String(http.StatusBadRequest, "bad request")
	}

	u, err := HandleRunScale(&cmd)

	if err != nil {
		return err
	}

	return scaleOutTemplate.ExecuteTemplate(context.Response(), "scale_out.html", u)
}
