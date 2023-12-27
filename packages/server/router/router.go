package router

import (
	. "se/controller"
	"se/controller/middleware"
)

func init() {
	login := Router.Group("/login")
	// patient := Router.Group("/patient").Use(middleware.IsPatient)
	// doctor := Router.Group("/doctor").Use(middleware.IsDoctor)
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
	}
}
