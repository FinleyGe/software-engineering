package model

import (
	. "se/db"
	"se/utility"
	"time"
)

type Patient struct {
	ID         uint64 `gorm:"primary_key"`
	DoctorID   uint64
	Name       string
	Password   string
	Phone      string
	Bed        string
	Birth      time.Time
	Gender     string
	VitalSigns []VitalSign
	InTime     time.Time `gorm:"required"`
	OutTime    time.Time
}

func (patient *Patient) CheckPassword(password string) bool {
	DB.Where("id = ?", patient.ID).First(patient)
	return patient.ID != 0 && utility.PasswordVerify(password, patient.Password)
}
