package controller

import (
	. "se/config"
	"se/db/model"
	. "se/utility"
	"strconv"

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
			if Config.Dev {
				c.SetCookie("token", token, 3600, "/", "localhost", false, false)
			} else {
				c.SetCookie("token", token, 3600, Config.Server.ApiRoute, Config.Server.Domain, false, false)
			}
			ResponseOK(c)
		}
		return
	}
	ResponseBadRequest(c)
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

func AddDoctor(c *gin.Context) {
	d := struct {
		Name         string `json:"name"`
		DepartmentID uint64 `json:"department_id"`
	}{}
	c.BindJSON(&d)
	doctor := model.Doctor{
		Name:         d.Name,
		DepartmentID: d.DepartmentID,
		Password: func() (pwd string) {
			pwd, _ = PasswordHash(Config.Server.DefaultPassword)
			return
		}(),
	}
	if doctor.Add() {
		ResponseOK(c)
	} else {
		ResponseServerError(c)
	}
}

type DoctorInfo struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	State          string `json:"state"`
	DepartmentID   uint64 `json:"department_id"`
	DepartmentName string `json:"department_name"`
	PatientCount   uint64 `json:"patient_count"`
}

func ShowDoctors(c *gin.Context) {
	doctors := model.GetAllDoctorsWithDepartment()
	doctorInfos := make([]DoctorInfo, len(doctors))
	for i, doctor := range doctors {
		doctorInfos[i] = DoctorInfo{
			ID:             doctor.ID,
			Name:           doctor.Name,
			DepartmentID:   doctor.DepartmentID,
			DepartmentName: doctor.Department.Name,
			PatientCount:   uint64(doctor.GetPatientsCount()),
			State:          doctor.State,
		}
		ResponseOKWithData(c, doctorInfos)
	}
}

func ShowDoctor(c *gin.Context) {
	idString := c.Param("id")
	doctor := model.GetDoctorByID(func() (id uint64) {
		id, _ = strconv.ParseUint(idString, 10, 64)
		return
	}())
	if doctor.ID == 0 {
		ResponseNotFound(c)
		return
	}
	ResponseOKWithData(c, doctor)
}

func ShowBeds(c *gin.Context) {
	beds := model.GetAllBeds()
	ResponseOKWithData(c, beds)
}

func ShowBed(c *gin.Context) {
	idString := c.Param("id")
	bed := model.GetBedByID(func() (id uint64) {
		id, _ = strconv.ParseUint(idString, 10, 64)
		return
	}())
	if bed.ID == 0 {
		ResponseNotFound(c)
		return
	}
	ResponseOKWithData(c, bed)
}

func AddBed(c *gin.Context) {
	b := struct {
		Number string `json:"number"`
	}{}

	c.BindJSON(&b)
	bed := model.Bed{
		BedNumber: b.Number,
	}
	if bed.Add() {
		ResponseOK(c)
	} else {
		ResponseServerError(c)
	}
}
