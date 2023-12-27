package controller

import (
	. "se/config"
	"se/db/model"
	. "se/utility"

	"github.com/gin-gonic/gin"
)

func DoctorLogin(c *gin.Context) {
	loginRequest := struct {
		ID       uint64 `json:"id"`
		Password string `json:"password"`
	}{}
	c.BindJSON(&loginRequest)
	doctor := model.Doctor{
		ID: loginRequest.ID,
	}
	if doctor.CheckPassword(loginRequest.Password) {
		token, err := GenerateToken(Data{
			ID:   doctor.ID,
			Role: "doctor",
		})
		if err != nil {
			ResponseServerError(c)
		} else {
			if Config.Dev {
				c.SetCookie("token", token, 3600, "/", "localhost", false, true)
			} else {
				c.SetCookie("token", token, 3600, Config.Server.ApiRoute, Config.Server.Domain, false, true)
			}
			ResponseOK(c)
		}
	}
	ResponseBadRequest(c)
}
