package router

import (
	. "se/controller"
	"se/controller/middleware"
)

func init() {
	login := Router.Group("/login")

	{
		login.POST("/patient", PatientLogin)
		login.POST("/doctor", DoctorLogin)
		login.POST("/admin", AdminLogin)
	}

	patient := Router.Group("/patient").Use(middleware.IsPatient)
	doctor := Router.Group("/doctor").Use(middleware.IsDoctor)
	admin := Router.Group("/admin").Use(middleware.IsAdmin)

}
