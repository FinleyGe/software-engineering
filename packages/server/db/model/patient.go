package model

import (
	. "se/db"
	"se/utility"
)

type Patient struct {
	ID       uint64       `gorm:"primary_key" json:"id"`
	DoctorID uint64       `gorm:"not null" json:"doctor_id"`
	Name     string       `gorm:"not null" json:"name"`
	Password string       `gorm:"not null" json:"-"`
	Phone    string       `gorm:"not null" json:"phone"`
	BedID    uint64       `gorm:"not null" json:"bed_id"`
	Birth    utility.Time `gorm:"not null" json:"birth"`
	Gender   string       `gorm:"not null" json:"gender"`
	InTime   utility.Time `json:"in_time"`
	OutTime  utility.Time `json:"out_time`
	State    string       `gorm:"not null" json:"state"`
	IsOut    bool         `gorm:"not null" json:"is_out"`
}

func (patient *Patient) CheckPassword(password string) bool {
	DB.Where("id = ?", patient.ID).First(patient)
	return patient.ID != 0 && utility.PasswordVerify(password, patient.Password)
}

func (patient *Patient) Add() bool {
	return DB.Create(patient).Error == nil
}

func GetAllPatient() []Patient {
	var patients []Patient
	DB.Find(&patients)
	return patients
}

func GetPatientByID(id uint64) Patient {
	var patient Patient
	DB.Where("id = ?", id).First(&patient)
	return patient
}

func (patient *Patient) Call() {
	DB.Model(patient).Update("state", "call")
}

func (p *Patient) GetDepartmentID() uint64 {
	if p.DoctorID == 0 {
		DB.Where("id = ?", p.ID).First(p)
	}
	doctor := Doctor{}
	DB.Where("id = ?", p.DoctorID).First(&doctor)
	return doctor.DepartmentID
}
