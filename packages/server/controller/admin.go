package controller

import (
	. "se/config"
	"se/db/model"
	"se/utility"
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
	}
	ResponseOKWithData(c, doctorInfos)
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
		Number       string `json:"number"`
		DepartmentID uint64 `json:"department_id"`
	}{}

	c.BindJSON(&b)
	bed := model.Bed{
		BedNumber:    b.Number,
		DepartmentID: b.DepartmentID,
	}
	if bed.Add() {
		ResponseOK(c)
	} else {
		ResponseServerError(c)
	}
}

func UpdateDepartment(c *gin.Context) {
	d := struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}{}
	c.BindJSON(&d)
	department := model.Department{
		ID:   d.ID,
		Name: d.Name,
	}
	if department.Update() {
		ResponseOK(c)
	} else {
		ResponseServerError(c)
	}
}

func DeleteDepartment(c *gin.Context) {
	idString := c.Query("id")
	department := model.Department{
		ID: func() (id uint64) {
			id, _ = strconv.ParseUint(idString, 10, 64)
			return
		}(),
	}
	if department.Delete() {
		ResponseOK(c)
	} else {
		ResponseServerError(c)
	}
}

func UpdateDoctorDepartment(c *gin.Context) {
	d := struct {
		ID           uint64 `json:"id"`
		DepartmentID uint64 `json:"department_id"`
	}{}
	c.BindJSON(&d)
	doctor := model.Doctor{
		ID:           d.ID,
		DepartmentID: d.DepartmentID,
	}
	if doctor.UpdateDepartment(d.DepartmentID) {
		ResponseOK(c)
	} else {
		ResponseServerError(c)
	}
}

func DeleteDoctor(c *gin.Context) {
	idString := c.Query("id")
	doctor := model.Doctor{
		ID: func() (id uint64) {
			id, _ = strconv.ParseUint(idString, 10, 64)
			return
		}(),
	}
	if doctor.Delete() {
		ResponseOK(c)
	} else {
		ResponseServerError(c)
	}
}

func DeleteBed(c *gin.Context) {
	idString := c.Query("id")
	bed := model.Bed{
		ID: func() (id uint64) {
			id, _ = strconv.ParseUint(idString, 10, 64)
			return
		}(),
	}
	if bed.Delete() {
		ResponseOK(c)
	} else {
		ResponseServerError(c)
	}
}

func DeletePatient(c *gin.Context) {
	idString := c.Query("id")
	patient := model.Patient{
		ID: func() (id uint64) {
			id, _ = strconv.ParseUint(idString, 10, 64)
			return
		}(),
	}
	if patient.Delete() {
		ResponseOK(c)
	} else {
		ResponseServerError(c)
	}
}

func DoctorLetPatientOut(c *gin.Context) {
	idString := c.Param("id")
	patient := model.GetPatientByID(func() (id uint64) {
		id, _ = strconv.ParseUint(idString, 10, 64)
		return
	}())
	if patient.LetOut() {
		ResponseOK(c)
	} else {
		ResponseServerError(c)
	}
}

func UpdatePatient(c *gin.Context) {
	p := struct {
		ID     uint64       `json:"id"`
		Name   string       `json:"name"`
		InTime utility.Time `json:"in_time"`
		State  string       `json:"state"`
		Gender string       `json:"state"`
		Phone  string       `json:"phone"`
	}{}
	c.BindJSON(&p)
	patient := model.GetPatientByID(p.ID)
	patient.Name = p.Name
	patient.InTime = p.InTime
	patient.State = p.State
	patient.Gender = p.Gender
	patient.Phone = p.Phone
	if patient.Update() {
		ResponseOK(c)
	} else {
		ResponseServerError(c)
	}
}
