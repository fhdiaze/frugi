package money

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	convertGetTemplate = template.Must(template.ParseFiles("templates/comps/frame.html", "templates/money/convert.html"))
	convertOutTemplate = template.Must(template.ParseFiles("templates/comps/frame.html", "templates/money/convert_out.html"))
)

func Route(group *echo.Group) {
	group.GET("/money.convert.get", handleGetConvert)
	group.POST("/money.convert.run", handleRunConvert)
}

func handleGetConvert(context echo.Context) error {
	return convertGetTemplate.ExecuteTemplate(context.Response(), "convert.html", nil)
}

func handleRunConvert(context echo.Context) error {
	var cmd RunConvertCmd
	err := context.Bind(&cmd)

	if err != nil {
		return context.String(http.StatusInternalServerError, "Internal Server Error")
	}

	u, err := HandleRunConvert(&cmd)

	if err != nil {
		return err
	}

	return convertOutTemplate.ExecuteTemplate(context.Response(), "convert_out.html", u)
}
