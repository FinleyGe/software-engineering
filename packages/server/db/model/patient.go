package model

import (
	. "se/db"
	"se/utility"

	"gorm.io/gorm"
)

type Patient struct {
	ID           uint64         `gorm:"primary_key" json:"id"`
	DoctorID     uint64         `gorm:"not null" json:"doctor_id"`
	Name         string         `gorm:"not null" json:"name"`
	Password     string         `gorm:"not null" json:"-"`
	Phone        string         `gorm:"not null" json:"phone"`
	BedID        uint64         `gorm:"not null" json:"bed_id"`
	Birth        utility.Time   `gorm:"not null" json:"birth"`
	Gender       string         `gorm:"not null" json:"gender"`
	InTime       utility.Time   `json:"in_time"`
	OutTime      utility.Time   `json:"out_time"`
	State        string         `gorm:"not null" json:"state"`
	IsOut        bool           `gorm:"not null" json:"is_out"`
	DeleteAt     gorm.DeletedAt `json:"-"`
	Logs         []Log          `json:"logs"`
	VitalSigns   []VitalSign    `json:"vital_signs"`
	DepartmentID uint64         `json:"department_id"`
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

func (p *Patient) GetDoctor() Doctor {
	doctor := Doctor{}
	DB.Where("id = ?", p.DoctorID).First(&doctor)
	return doctor
}

func (p *Patient) GetBed() Bed {
	bed := Bed{}
	DB.Where("id = ?", p.BedID).First(&bed)
	return bed
}

func (p *Patient) GetDepartment() Department {
	department := Department{}
	DB.Where("id = ?", p.GetDepartmentID()).First(&department)
	return department
}

func (p *Patient) Delete() bool {
	return DB.Delete(p).Error == nil
}

func (p *Patient) Update() bool {
	return DB.Model(p).Updates(p).Error == nil
}

func GetPatientsByDoctorID(id uint64) []Patient {
	var patients []Patient
	DB.Where("doctor_id = ?", id).Find(&patients)
	return patients
}

func (p *Patient) GetLogs() []Log {
	var logs []Log
	DB.Where("patient_id = ?", p.ID).Find(&logs)
	return logs
}

func (p *Patient) GetVitalSigns() []VitalSign {
	var vitalSigns []VitalSign
	DB.Where("patient_id = ?", p.ID).Find(&vitalSigns)
	return vitalSigns
}

func (p *Patient) LetOut() bool {
	p.State = "out"
	p.IsOut = true
	p.OutTime = utility.TimeNow()
	bed := p.GetBed()
	bed.Occupied = false
	bed.SetEmpty()
	return DB.Save(p).Error == nil
}

func GetPatientByBedID(id uint64) Patient {
	var patient Patient
	DB.Where("bed_id = ?", id).First(&patient)
	return patient
}
