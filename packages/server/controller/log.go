package controller

import (
	"github.com/gin-gonic/gin"
	"se/db/model"
	. "se/utility"
	"strconv"
)

// doctor.GET("/logs", DoctorGetAllLogs)
// doctor.GET("/logs/:patiend_id", DoctorGetPatientLogs)
// doctor.GET("/log/:id", DoctorGetLog)
// doctor.PUT("/log/:id", DoctorUpdateLog)
// doctor.DELETE("/log/:id", DoctorDeleteLog)
// doctor.POST("/log", DoctorAddLog)

func DoctorGetAllLogs(c *gin.Context) {
	doctor_id := c.MustGet("id").(uint64)
	logs := model.GetAllLogsByDoctorID(doctor_id)
	ResponseOKWithData(c, logs)
}

func DoctorGetPatientLogs(c *gin.Context) {
	patient_id, _ := strconv.ParseUint(c.Param("patient_id"), 10, 64)
	logs := model.GetAllLogsByPatientID(patient_id)
	ResponseOKWithData(c, logs)
}

func DoctorGetLog(c *gin.Context) {
	log_id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	log := model.GetLogByID(log_id)
	ResponseOKWithData(c, log)
}

func DoctorUpdateLog(c *gin.Context) {
	type Request struct {
		ID      uint64 `json:"id"`
		Content string `json:"content"`
	}
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		ResponseBadRequest(c)
		return
	}
	log := model.GetLogByID(request.ID)
	if (log == model.Log{}) {
		ResponseNotFound(c)
		return
	}
	log.Content = request.Content
	log.Update()
}

func DoctorDeleteLog(c *gin.Context) {
	log_id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	log := model.GetLogByID(log_id)
	if (log == model.Log{}) {
		ResponseNotFound(c)
		return
	}
	log.Delete()
}

func DoctorAddLog(c *gin.Context) {
	type Request struct {
		PatientID uint64 `json:"patient_id"`
		Content   string `json:"content"`
	}
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		ResponseBadRequest(c)
		return
	}
	doctor_id := c.MustGet("id").(uint64)
	log := model.Log{
		DoctorID:  doctor_id,
		PatientID: request.PatientID,
		Content:   request.Content,
	}
	log.Add()
	ResponseOK(c)
}
