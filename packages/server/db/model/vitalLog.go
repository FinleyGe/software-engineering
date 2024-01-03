package model

import (
	. "se/db"
	"se/utility"
)

type VitalSign struct {
	ID            uint64       `gorm:"primary_key" json:"id"`
	BedID         uint64       `json:"bed_id"`
	PatientID     uint64       `json:"patient_id"`
	Time          utility.Time `json:"time"`
	Temperature   float64      `json:"temperature"`
	HeartRate     uint64       `json:"heart_rate"`
	BloodPressure string       `json:"blood_pressure"`
	BreathingRate uint64       `json:"breathing_rate"`
	BloodOxygen   float64      `json:"blood_oxygen"`
	Sense         string       `json:"sense"`
}

func (vitalSign *VitalSign) Add() error {
	p := GetPatientByBedID(vitalSign.BedID)
	vitalSign.PatientID = p.ID
	return DB.Create(vitalSign).Error
}

func GetVitalSignsByBedID(bedID uint64) ([]VitalSign, error) {
	var vitalSigns []VitalSign
	err := DB.Where("bed_id = ?", bedID).Find(&vitalSigns).Error
	return vitalSigns, err
}

func GetVitalSignsByBedIDAndTime(bedID uint64, start utility.Time, end utility.Time) ([]VitalSign, error) {
	var vitalSigns []VitalSign
	err := DB.Where("bed_id = ? AND time BETWEEN ? AND ?", bedID, start, end).Find(&vitalSigns).Error
	return vitalSigns, err
}
