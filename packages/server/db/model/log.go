package model

import (
	. "se/db"
	"se/utility"
)

type Log struct {
	ID        uint64       `gorm:"primary_key" json:"id"`
	Content   string       `gorm:"type:text" json:"content"`
	PatientID uint64       `gorm:"index" json:"patient_id"`
	DoctorID  uint64       `gorm:"index" json:"doctor_id"`
	Time      utility.Time `json:"time"`
	IsDeleted bool         `gorm:"default:false" json:"is_deleted"`
}

func (l *Log) Add() error {
	return DB.Create(l).Error
}

func GetAllLogsByPatientID(patient_id uint64) []Log {
	var logs []Log
	DB.Where("patient_id = ?", patient_id).Find(&logs)
	return logs
}

func GetLogByID(log_id uint64) Log {
	var log Log
	DB.Where("id = ?", log_id).First(&log)
	return log
}

func (l *Log) Delete() error {
	l.IsDeleted = true
	return DB.Save(l).Error
}

func (l *Log) Update() error {
	return DB.Save(l).Error
}

func GetAllLogsByDoctorID(doctor_id uint64) []Log {
	var logs []Log
	DB.Where("doctor_id = ?", doctor_id).Find(&logs)
	return logs
}
