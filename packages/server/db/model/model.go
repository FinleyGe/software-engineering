package model

import (
	"time"
)

type Department struct {
	ID   uint64 `gorm:"primary_key"`
	Name string
}

type VitalSign struct {
	ID            uint64 `gorm:"primary_key"`
	PatientID     uint64
	Time          time.Time
	Temperature   float64
	HeartRate     uint64
	BloodPressure string
	BreathingRate uint64
	BloodOxygen   uint64
	Sense         string
}
