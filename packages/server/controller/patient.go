package controller

import (
	"log"
	. "se/config"
	"se/db/model"
	"se/grpc"
	. "se/utility"

	"github.com/gin-gonic/gin"
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
				c.SetCookie("token", token, 3600, "/", "localhost", false, false)
			} else {
				c.SetCookie("token", token, 3600, Config.Server.ApiRoute, Config.Server.Domain, false, false)
			}
			ResponseOK(c)
		}
	} else {
		ResponseBadRequest(c)
	}
}

func Call(c *gin.Context) {
	patient_id := c.MustGet("id").(uint64)
	role := c.MustGet("role").(string)
	if role != "patient" {
		ResponseForbidden(c)
		return
	}
	err := grpc.Call(patient_id)
	if err != nil {
		log.Println(err)
		ResponseServerError(c)
	} else {
		ResponseOK(c)
	}
}
