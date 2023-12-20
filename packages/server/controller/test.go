package controller

import (
	. "Go-Server-Scaffold/utility"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	ResponseOK(c, "OK", nil)
}
