package controller

import (
	"se/db/model"
	. "se/utility"

	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	loginRequest := struct {
		ID       uint64 `json:"id"`
		Password string `json:"password"`
	}{}
	c.BindJSON(&loginRequest)
	admin := model.Admin{
		ID: loginRequest.ID,
	}
	if admin.CheckPassword(loginRequest.Password) {
		token, err := GenerateToken(Data{
			ID:   admin.ID,
			Role: "admin",
		})
		if err != nil {
			ResponseServerError(c)
		} else {
			ResponseOKWithData(c, gin.H{
				"token": token,
			})
		}
	}
}
