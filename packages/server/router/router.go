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
		admin.POST("/department", AddDepartment)
		admin.GET("/department", ShowDepartments)
		admin.POST("/doctor", AddDoctor)
		admin.GET("/doctors", ShowDoctors)
		admin.GET("/doctor/:id", ShowDoctor)
		admin.GET("/beds", ShowBeds)
		admin.POST("/bed", AddBed)
		admin.GET("/bed/:id", ShowBed)
		admin.POST("/patient", AddPatient)
		admin.GET("/patients", ShowPatients)
		admin.GET("/patient/:id", ShowPatient)
		admin.GET("/vital_signs/:bed_id", ShowVitalSignsWithTimeFilter)
		// admin.GET("/call", Call)
	}
	{
		doctor.GET("/patients", ShowPatients)
		doctor.GET("/patient/:id", ShowPatient)
		doctor.POST("/patients", AddPatient)
		doctor.GET("/beds", ShowBeds)
		doctor.GET("/bed/:id", ShowBed)
		doctor.GET("/vital_signs/:bed_id", ShowVitalSignsWithTimeFilter)
	}
	{
		patient.GET("/call", Call)
	}
}
