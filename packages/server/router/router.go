package router

import (
	. "se/controller"
	"se/controller/middleware"
)

func init() {
	login := Router.Group("/login")
	patient := Router.Group("/patient").Use(middleware.IsPatient)
	doctor := Router.Group("/doctor").Use(middleware.IsDoctor)
	admin := Router.Group("/admin").Use(middleware.IsAdmin)
	{
		login.POST("/patient", PatientLogin)
		login.POST("/doctor", DoctorLogin)
		login.POST("/admin", AdminLogin)
	}
	{
		admin.PUT("/department", UpdateDepartment)
		admin.POST("/department", AddDepartment)
		admin.GET("/department", ShowDepartments)
		admin.DELETE("/department", DeleteDepartment)

		admin.POST("/doctor", AddDoctor)
		admin.GET("/doctors", ShowDoctors)
		admin.GET("/doctor/:id", ShowDoctor)
		admin.PUT("/doctor", UpdateDoctorDepartment)
		admin.DELETE("/doctor", DeleteDoctor)

		admin.GET("/beds", ShowBeds)
		admin.POST("/bed", AddBed)
		admin.DELETE("/bed", DeleteBed)
		// admin.GET("/bed/:id", ShowBed)
		admin.POST("/patient", AddPatient)
		admin.GET("/patients", ShowPatients)
		admin.GET("/patient/:id", ShowPatient)
		admin.DELETE("/patient", DeletePatient)
		admin.PUT("/patient", UpdatePatient)
		admin.GET("/vital_signs/:bed_id", ShowVitalSignsWithTimeFilter)
		// admin.GET("/call", Call)
	}
	{
		doctor.GET("/patients", ShowDoctorPatients)
		doctor.GET("/patient/:id", ShowPatient)
		doctor.POST("/patient", DoctorAddPatient)
		doctor.GET("/beds", ShowBeds)
		doctor.GET("/bed/:id", ShowBed)
		doctor.GET("/vital_signs/:bed_id", ShowVitalSignsWithTimeFilter)
		doctor.DELETE("/patient/:id", DoctorLetPatientOut)

		doctor.GET("/logs", DoctorGetAllLogs)
		doctor.GET("/logs/:patiend_id", DoctorGetPatientLogs)
		doctor.GET("/log/:id", DoctorGetLog)
		doctor.PUT("/log/:id", DoctorUpdateLog)
		doctor.DELETE("/log/:id", DoctorDeleteLog)
		doctor.POST("/log", DoctorAddLog)
	}
	{
		patient.GET("/call", Call)
		patient.GET("/info", PatientShowInfo)
		patient.GET("/vital_signs", PatientShowVitalSignsWithTimeFilter)
	}
}
