package router

import (
	"se/controller/middleware"
	"se/controller/ws"
)

func init() {
	Router.GET("/ws",  middleware.IsDoctor, ws.HandleWSConnection)
}
