package model

import (
	. "se/db"
	"se/utility"
	"time"
)

type Doctor struct {
	ID           uint64 `gorm:"primary_key"`
	Birth        time.Time
	Name         string
	Password     string
	DepartmentID uint64
	Department   Department
	Patients     []Patient `gorm:"many2many:doctor_patients;"`
}

func (doctor *Doctor) CheckPassword(password string) bool {
	DB.Where("id = ?", doctor.ID).First(doctor)
	return doctor.ID != 0 && utility.PasswordVerify(password, doctor.Password)
}
