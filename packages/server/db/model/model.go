package model

import (
	"time"
)

type VitalSign struct {
	ID            uint64 `gorm:"primary_key"`
	PatientID     uint64
	Time          time.Time
	Temperature   float64
	HeartRate     uint64
	BloodPressure float64
	BreathingRate uint64
	BloodOxygen   float64
	Sense         string
}
