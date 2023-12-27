package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	. "se/config"
)

func Cors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	if Config.Dev {
		corsConfig.AllowAllOrigins = true
	} else {

		// TODO: add AllowHeaders
		/* cors.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization") */
		corsConfig.AllowOrigins = Config.Server.AllowOrigins
	}
	return cors.New(corsConfig)
}
