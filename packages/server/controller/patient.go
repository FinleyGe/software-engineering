package controller

import (
	"se/db/model"
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
			ResponseOKWithData(c, gin.H{"token": token})
		}
	} else {
		ResponseUnauthorized(c)
	}
}
