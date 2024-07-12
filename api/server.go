package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {
	RouteHome(engine)
	RouteCmd(engine)
}

func Static(engine *gin.Engine) {
	engine.StaticFS("/assets", http.Dir("assets"))
}
