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
	doctor := model.GetDoctorByID(p.DoctorID)
	patient := model.Patient{
		Name:         p.Name,
		DoctorID:     p.DoctorID,
		BedID:        p.BedID,
		IsOut:        false,
		State:        p.State,
		Gender:       p.Gender,
		Phone:        p.Phone,
		Birth:        p.Birth,
		InTime:       p.InTime,
		Password:     utility.GetDefaultPassword(),
		DepartmentID: doctor.DepartmentID,
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

type apiDoctor struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
type apiBed struct {
	ID     uint64 `json:"id"`
	Number string `json:"number"`
}

type apiPatient struct {
	ID      uint64       `json:"id"`
	Name    string       `json:"name"`
	InTime  utility.Time `json:"in_time"`
	OutTime utility.Time `json:"out_time"`
	State   string       `json:"state"`
	Doctor  apiDoctor    `json:"doctor"`
	Bed     apiBed       `json:"bed"`
}

type apiPatientDetail struct {
	ID      uint64       `json:"id"`
	Name    string       `json:"name"`
	InTime  utility.Time `json:"in_time"`
	OutTime utility.Time `json:"out_time"`
	State   string       `json:"state"`
	Gender  string       `json:"gender"`
	Birth   utility.Time `json:"birth"`
	Phone   string       `json:"phone"`
	Doctor  apiDoctor    `json:"doctor"`
	Bed     apiBed       `json:"bed"`
}

func ShowPatients(c *gin.Context) {
	patients := model.GetAllPatient()
	var apiPatients []apiPatient
	for _, patient := range patients {
		apiPatient := apiPatient{
			ID:     patient.ID,
			Name:   patient.Name,
			InTime: patient.InTime,
			OutTime: func() utility.Time {
				if patient.IsOut {
					return patient.OutTime
				}
				return utility.Time{}
			}(),
			State: patient.State,
			Doctor: func() apiDoctor {
				doctor := model.GetDoctorByID(patient.DoctorID)
				return apiDoctor{
					ID:   doctor.ID,
					Name: doctor.Name,
				}
			}(),
			Bed: func() apiBed {
				bed := model.GetBedByID(patient.BedID)
				return apiBed{
					ID:     bed.ID,
					Number: bed.BedNumber,
				}
			}(),
		}
		apiPatients = append(apiPatients, apiPatient)
	}
	ResponseOKWithData(c, apiPatients)
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
	var apiPatient apiPatientDetail
	apiPatient.ID = patient.ID
	apiPatient.Name = patient.Name
	apiPatient.InTime = patient.InTime
	apiPatient.OutTime = func() utility.Time {
		if patient.IsOut {
			return patient.OutTime
		}
		return utility.Time{}
	}()
	apiPatient.State = patient.State
	apiPatient.Doctor = func() apiDoctor {
		doctor := model.GetDoctorByID(patient.DoctorID)
		return apiDoctor{
			ID:   doctor.ID,
			Name: doctor.Name,
		}
	}()
	apiPatient.Bed = func() apiBed {
		bed := model.GetBedByID(patient.BedID)
		return apiBed{
			ID:     bed.ID,
			Number: bed.BedNumber,
		}
	}()
	apiPatient.Birth = patient.Birth
	apiPatient.Phone = patient.Phone
	apiPatient.Gender = patient.Gender
	ResponseOKWithData(c, apiPatient)
	// ResponseOKWithData(c, patient)
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

func DoctorAddPatient(c *gin.Context) {

	patient := struct {
		Name   string       `json:"name"`
		State  string       `json:"state"`
		Gender string       `json:"gender"`
		Bitrh  utility.Time `json:"birth"`
		Phone  string       `json:"phone"`
		BedID  uint64       `json:"bed_id"`
	}{}

	if err := c.BindJSON(&patient); err != nil {
		ResponseBadRequest(c)
		return
	}
	doctorID := c.MustGet("id").(uint64)
	doctor := model.GetDoctorByID(doctorID)
	p := model.Patient{
		Name:         patient.Name,
		State:        patient.State,
		Phone:        patient.Phone,
		Gender:       patient.Gender,
		Birth:        patient.Bitrh,
		InTime:       utility.TimeNow(),
		BedID:        patient.BedID,
		DoctorID:     doctorID,
		Password:     utility.GetDefaultPassword(),
		DepartmentID: doctor.DepartmentID,
	}
	if p.Add() {
		bed := model.GetBedByID(patient.BedID)
		bed.Occupied = true
		bed.Update()
		ResponseOK(c)
	} else {
		ResponseBadRequest(c)
	}
}

type apiDoctorPatient struct {
	ID      uint64       `json:"id"`
	Name    string       `json:"name"`
	InTime  utility.Time `json:"in_time"`
	OutTime utility.Time `json:"out_time"`
	State   string       `json:"state"`
	Bed     apiBed       `json:"bed"`
}

func ShowDoctorPatients(c *gin.Context) {
	id := func() uint64 {
		id, _ := c.Get("id")
		return id.(uint64)
	}()
	patients := model.GetPatientsByDoctorID(id)
	apiPatients := []apiDoctorPatient{}
	for _, patient := range patients {
		apiPatient := apiDoctorPatient{
			ID:     patient.ID,
			Name:   patient.Name,
			InTime: patient.InTime,
			OutTime: func() utility.Time {
				if patient.IsOut {
					return patient.OutTime
				}
				return utility.Time{}
			}(),
			State: patient.State,
			Bed: func() apiBed {
				bed := model.GetBedByID(patient.BedID)
				return apiBed{
					ID:     bed.ID,
					Number: bed.BedNumber,
				}
			}(),
		}
		apiPatients = append(apiPatients, apiPatient)
	}
	ResponseOKWithData(c, apiPatients)
}
