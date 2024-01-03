package model

import (
	. "se/db"

	"gorm.io/gorm"
)

type Bed struct {
	ID           uint64         `json:"id" gorm:"primary_key"`
	BedNumber    string         `json:"bed_number" gorm:"uniqueIndex"`
	Occupied     bool           `json:"occupied" gorm:"default:false"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	DepartmentID uint64         `json:"department_id"`
	Department   Department     `json:"-"`
}

func (bed *Bed) Add() bool {
	return DB.Create(bed).Error == nil
}

func GetAllBeds() []Bed {
	var beds []Bed
	DB.Find(&beds)
	return beds
}

func GetBedByID(id uint64) Bed {
	var bed Bed
	DB.First(&bed, id)
	return bed
}

func GetBedByBedNumber(bedNumber string) Bed {
	var bed Bed
	DB.Where("bed_number = ?", bedNumber).First(&bed)
	return bed
}

func (bed *Bed) Update() bool {
	return DB.Save(bed).Error == nil
}

func (bed *Bed) Delete() bool {
	return DB.Delete(bed).Error == nil
}

func (bed *Bed) SetOccupied() bool {
	bed.Occupied = true
	return bed.Update()
}

func (bed *Bed) SetEmpty() bool {
	bed.Occupied = false
	return bed.Update()
}
