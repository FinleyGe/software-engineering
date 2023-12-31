package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	. "se/config"
	"se/db/model"
	"se/utility"
	. "se/utility"
	"strconv"
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

func AddPatient(c *gin.Context) {
	p := struct {
		Name     string       `json:"name"`
		DoctorID uint64       `json:"doctor_id"`
		BedID    uint64       `json:"bed_id"`
		InTime   utility.Time `json:"in_time"`
		State    string       `json:"state"`
		Gender   string       `json:"gender"`
		Phone    string       `json:"phone"`
		Birth    utility.Time `json:"birth"`
	}{}
	c.BindJSON(&p)
	patient := model.Patient{
		Name:     p.Name,
		DoctorID: p.DoctorID,
		BedID:    p.BedID,
		IsOut:    false,
		State:    p.State,
		Gender:   p.Gender,
		Phone:    p.Phone,
		Birth:    p.Birth,
		InTime:   p.InTime,
		Password: utility.GetDefaultPassword(),
	}
	if !patient.Add() {
		ResponseServerError(c)
		return
	}
	bed := model.GetBedByID(p.BedID)
	bed.Occupied = true
	bed.Update()
	ResponseOK(c)
}

func ShowPatients(c *gin.Context) {
	patients := model.GetAllPatient()
	ResponseOKWithData(c, patients)
}

func ShowPatient(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		ResponseBadRequest(c)
		return
	}
	patient := model.GetPatientByID(
		func() uint64 {
			i, _ := strconv.ParseUint(id, 10, 64)
			return i
		}(),
	)
	ResponseOKWithData(c, patient)
}

func ShowVitalSigns(c *gin.Context) {
	bedID := c.Param("bed_id")
	if bedID == "" {
		ResponseBadRequest(c)
		return
	}
	vitalSigns, err := model.GetVitalSignsByBedID(
		func() uint64 {
			i, _ := strconv.ParseUint(bedID, 10, 64)
			return i
		}(),
	)
	if err != nil {
		ResponseServerError(c)
		return
	}
	ResponseOKWithData(c, vitalSigns)
}

func ShowVitalSignsWithTimeFilter(c *gin.Context) {
	timeStart := utility.AsTime(c.Query("time_start"))
	timeUntil := utility.AsTime(c.Query("time_until"))
	log.Println(timeStart)
	log.Println(timeUntil)
	if timeStart.IsZero() {
		timeStart = utility.AsTime("0001-01-01 00:00:00")
	}
	if timeUntil.IsZero() {
		timeUntil = utility.AsTime("2100-01-01 00:00:00")
	}
	bedID := c.Param("bed_id")
	if bedID == "" {
		ResponseBadRequest(c)
		return
	}
	vitalSigns, err := model.GetVitalSignsByBedIDAndTime(
		func() uint64 {
			i, _ := strconv.ParseUint(bedID, 10, 64)
			return i
		}(),
		timeStart,
		timeUntil,
	)
	if err != nil {
		ResponseServerError(c)
		return
	}
	ResponseOKWithData(c, vitalSigns)

}
