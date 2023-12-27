package middleware

import (
	"github.com/gin-gonic/gin"
	. "se/utility"
)

func IsDoctor(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		ResponseUnauthorized(c)
		c.Abort()
		return
	}
	data, err := ParseToken(token)
	if data.Role != "doctor" {
		ResponseUnauthorized(c)
		c.Abort()
		return
	}
	if err != nil {
		ResponseUnauthorized(c)
		c.Abort()
		return
	}
	c.Set("id", data.ID)
	c.Set("role", data.Role)
	c.Next()
}

func IsPatient(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		ResponseUnauthorized(c)
		c.Abort()
		return
	}
	data, err := ParseToken(token)
	if data.Role != "patient" {
		ResponseUnauthorized(c)
		c.Abort()
		return
	}
	if err != nil {
		ResponseUnauthorized(c)
		c.Abort()
		return
	}
	c.Set("id", data.ID)
	c.Set("role", data.Role)
	c.Next()
}

func IsAdmin(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		ResponseUnauthorized(c)
		c.Abort()
		return
	}
	data, err := ParseToken(token)
	if data.Role != "admin" {
		ResponseUnauthorized(c)
		c.Abort()
		return
	}
	if err != nil {
		ResponseUnauthorized(c)
		c.Abort()
		return
	}
	c.Set("id", data.ID)
	c.Set("role", data.Role)
	c.Next()
}
