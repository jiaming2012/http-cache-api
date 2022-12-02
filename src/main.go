package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jiaming2012/http-cache-api/src/routes"
)

func main() {
	routes.Router = gin.Default()

	routes.Init()

	routes.Router.Run()
}
