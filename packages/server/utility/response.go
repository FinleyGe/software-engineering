package utility

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, code int, message string, data interface{}) {
	if data == nil {
		c.JSON(code, gin.H{
			"message": message,
		})
	} else {
		c.JSON(code, gin.H{
			"message": message,
			"data":    data,
		})
	}
}

func ResponseOK(c *gin.Context) {
	Response(c, http.StatusOK, "OK", nil)
}

func ResponseOKWithData(c *gin.Context, data interface{}) {
	Response(c, http.StatusOK, "OK", data)
}

func ResponseBadRequest(c *gin.Context) {
	Response(c, http.StatusBadRequest, "Bad Request", nil)
}

func ResponseUnauthorized(c *gin.Context) {
	Response(c, http.StatusUnauthorized, "Unauthorized", nil)
}

func ResponseForbidden(c *gin.Context) {
	Response(c, http.StatusForbidden, "Forbidden", nil)
}

func ResponseServerError(c *gin.Context) {
	Response(c, http.StatusInternalServerError, "Server Error", nil)
}
