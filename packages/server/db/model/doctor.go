package model

import (
	. "se/db"
	"se/utility"
	"time"
)

type Doctor struct {
	ID           uint64 `gorm:"primary_key"`
	Birth        time.Time
	Name         string `gorm:"not null"`
	Password     string `gorm:"not null" json:"-"`
	DepartmentID uint64 `gorm:"not null"`
	Department   Department
	Patients     []Patient `gorm:"many2many:doctor_patients;"`
}

func (doctor *Doctor) CheckPassword(password string) bool {
	DB.Where("id = ?", doctor.ID).First(doctor)
	return doctor.ID != 0 && utility.PasswordVerify(password, doctor.Password)
}

func (doctor *Doctor) Add() bool {
	return DB.Create(doctor).Error == nil
}

func GetAllDoctors() []Doctor {
	doctors := []Doctor{}
	DB.Find(&doctors)
	return doctors
}

func GetDoctorByID(id uint64) Doctor {
	doctor := Doctor{}
	DB.Where("id = ?", id).First(&doctor)
	return doctor
}

func GetDoctorByName(name string) Doctor {
	doctor := Doctor{}
	DB.Where("name = ?", name).First(&doctor)
	return doctor
}
