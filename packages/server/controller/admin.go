package controller

import (
	"github.com/gin-gonic/gin"
	. "se/config"
	"se/db/model"
	. "se/utility"
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
			if Config.Dev {
				c.SetCookie("token", token, 3600, "/", "localhost", false, true)
			} else {
				c.SetCookie("token", token, 3600, Config.Server.ApiRoute, Config.Server.Domain, false, true)
			}
			ResponseOK(c)
		}
	}
}

func AddDepartment(c *gin.Context) {
	d := struct {
		Name string `json:"name"`
	}{}
	c.BindJSON(&d)
	department := model.Department{
		Name: d.Name,
	}
	if department.Add() {
		ResponseOK(c)
	} else {
		ResponseServerError(c)
	}
}

func ShowDepartments(c *gin.Context) {
	departments := model.GetAllDepartment()
	ResponseOKWithData(c, departments)
}
