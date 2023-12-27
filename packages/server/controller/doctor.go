package controller

import (
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
			ResponseOKWithData(c, gin.H{
				"token": token,
			})
		}
	}
	ResponseBadRequest(c)
}
