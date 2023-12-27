package controller

import (
	"github.com/gin-gonic/gin"
	. "se/config"
	"se/db/model"
	. "se/utility"
)

func PatientLogin(c *gin.Context) {
	loginRequest := struct {
		ID       uint64 `json:"id"`
		Password string `json:"password"`
	}{}
	c.BindJSON(&loginRequest)
	patient := model.Patient{}
	patient.ID = loginRequest.ID
	if patient.CheckPassword(loginRequest.Password) {
		token, err := GenerateToken(
			Data{
				ID:   patient.ID,
				Role: "patient",
			},
		)
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
	} else {
		ResponseUnauthorized(c)
	}
}
