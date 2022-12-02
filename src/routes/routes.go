package routes

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Init() {
	Router.DELETE("/cache", CreateDeleteCachedEntityEvent)
}
