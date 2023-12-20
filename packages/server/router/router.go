package router

import (
	. "Go-Server-Scaffold/controller"
)

func init() {
	Router.GET("/test", Test)
}
