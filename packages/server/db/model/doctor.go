package model

import (
	. "se/db"
	"se/utility"
)

type Doctor struct {
	ID           uint64       `gorm:"primary_key" json:"id"`
	Birth        utility.Time `json:"birth"`
	Name         string       `gorm:"not null" json:"name"`
	Password     string       `gorm:"not null" json:"-"`
	State        string       `gorm:"not null" json:"state"`
	DepartmentID uint64       `gorm:"not null" json:"department_id"`
	Department   Department   `json:"department"`
	Patients     []Patient    `gorm:"many2many:doctor_patients;" json:"patients"`
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

func GetAllDoctorsWithDepartment() []Doctor {
	doctors := []Doctor{}
	DB.Preload("Department").Find(&doctors)
	return doctors
}

func GetAllDoctorsWithPatients() []Doctor {
	doctors := []Doctor{}
	DB.Preload("Patients").Find(&doctors)
	return doctors
}

func GetAllDoctorsWithDepartmentAndPatients() []Doctor {
	doctors := []Doctor{}
	DB.Preload("Department").Preload("Patients").Find(&doctors)
	return doctors
}

func GetDoctorByID(id uint64) Doctor {
	doctor := Doctor{}
	DB.Where("id = ?", id).Preload("Department").Preload("Patients").First(&doctor)
	return doctor
}

func GetDoctorByName(name string) Doctor {
	doctor := Doctor{}
	DB.Where("name = ?", name).First(&doctor)
	return doctor
}

func (doctor *Doctor) GetDepartment() Department {
	DB.Where("id = ?", doctor.DepartmentID).First(&doctor.Department)
	return doctor.Department
}

func (doctor *Doctor) GetPatients() []Patient {
	DB.Model(doctor).Association("Patients").Find(&doctor.Patients)
	return doctor.Patients
}

func (doctor *Doctor) GetPatientsCount() int64 {
	return DB.Model(doctor).Association("Patients").Count()
}
