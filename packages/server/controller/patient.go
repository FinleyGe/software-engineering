package controller

import (
	"log"
	. "se/config"
	"se/db/model"
	"se/grpc"
	"se/utility"
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

func PatientShowInfo(c *gin.Context) {
	patient_id := c.MustGet("id").(uint64)
	patient := model.GetPatientByID(patient_id)
	apiPatient := apiPatientDetail{
		ID:      patient.ID,
		Name:    patient.Name,
		State:   patient.State,
		InTime:  patient.InTime,
		OutTime: patient.OutTime,
		Gender:  patient.Gender,
		Birth:   patient.Birth,
		Phone:   patient.Phone,
		Doctor: apiDoctor{
			ID:   patient.DoctorID,
			Name: patient.GetDoctor().Name,
		},
		Bed: apiBed{
			ID:     patient.BedID,
			Number: patient.GetBed().BedNumber,
		},
	}
	ResponseOKWithData(c, apiPatient)
}

func PatientShowVitalSignsWithTimeFilter(c *gin.Context) {
	id := c.MustGet("id").(uint64)
	patient := model.GetPatientByID(id)
	bedID := patient.BedID
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
	if bedID == 0 {
		ResponseBadRequest(c)
		return
	}
	vitalSigns, err := model.GetVitalSignsByBedIDAndTime(
		bedID,
		timeStart,
		timeUntil,
	)
	if err != nil {
		ResponseServerError(c)
		return
	}
	ResponseOKWithData(c, vitalSigns)
}
