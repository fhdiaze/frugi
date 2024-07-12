package main

import (
	"github.com/fhdiaze/frugi/api"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	api.Route(engine)
	api.Static(engine)

	engine.Run()
}
