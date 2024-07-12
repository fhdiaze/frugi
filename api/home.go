package api

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func RouteHome(router *gin.Engine) {
	router.GET("/", getHome)
}

func getHome(context *gin.Context) {
	homeTemplate := template.Must(template.ParseFiles("templates/base.html", "templates/home.html"))
	homeTemplate.Execute(context.Writer, nil)
}
