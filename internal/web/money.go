package web

import (
	"html/template"
	"net/http"

	"github.com/fhdiaze/frugi/internal/commands"
	"github.com/fhdiaze/frugi/resources"
	"github.com/labstack/echo/v4"
)

var (
	convertGetTemplate  = template.Must(template.ParseFS(resources.TemplatesFS, "templates/comps/frame.html", "templates/money/convert.html"))
	convertOutTemplate  = template.Must(template.ParseFS(resources.TemplatesFS, "templates/comps/frame.html", "templates/money/convert_out.html"))
	compoundGetTemplate = template.Must(template.ParseFS(resources.TemplatesFS, "templates/comps/frame.html", "templates/money/compound.html"))
	compoundOutTemplate = template.Must(template.ParseFS(resources.TemplatesFS, "templates/comps/frame.html", "templates/money/compound_out.html"))
)

func RouteMoney(group *echo.Group) {
	group.GET("/money.convert.get", handleGetConvert)
	group.POST("/money.convert.run", handleRunConvert)
	group.GET("/money.compound.get", handleGetCompound)
	group.POST("/money.compound.run", handleRunCompound)
}

func handleGetConvert(context echo.Context) error {
	return convertGetTemplate.ExecuteTemplate(context.Response(), "convert.html", nil)
}

func handleRunConvert(context echo.Context) error {
	var cmd commands.RunConvertCmd
	err := context.Bind(&cmd)

	if err != nil {
		return context.String(http.StatusInternalServerError, "Internal Server Error")
	}

	u, err := commands.HandleRunConvert(&cmd)

	if err != nil {
		return err
	}

	return convertOutTemplate.ExecuteTemplate(context.Response(), "convert_out.html", u)
}

func handleGetCompound(context echo.Context) error {
	return compoundGetTemplate.ExecuteTemplate(context.Response(), "compound.html", commands.AllFrequencyNames())
}

func handleRunCompound(context echo.Context) error {
	var cmd commands.RunCompoundCmd
	if err := context.Bind(&cmd); err != nil {
		return err
	}

	result := commands.HandleRunCompound(&cmd)

	return compoundOutTemplate.ExecuteTemplate(context.Response(), "compound_out.html", result.ToFloat64())
}
