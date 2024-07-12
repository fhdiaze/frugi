package api

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

type RunCmd struct {
	Cmd string `form:"cmd"`
}

func RouteCmd(router *gin.Engine) {
	router.POST("/cmd/run", run)
}

func run(context *gin.Context) {
	var runCmd RunCmd
	context.Bind(&runCmd)
	template := template.Must(template.ParseFiles("templates/output.html"))
	template.Execute(context.Writer, runCmd.Cmd)
}
