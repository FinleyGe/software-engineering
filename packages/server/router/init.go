package router

import (
	"github.com/gin-gonic/gin"
	. "se/controller/middleware"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.Use(Cors())
}
