package model

import (
	. "se/db"
)

type Bed struct {
	ID        uint64 `json:"id" gorm:"primary_key"`
	BedNumber string `json:"bed_number" gorm:"type:varchar(255);not null;unique_index"`
	Occupied  bool   `json:"occupied" gorm:"default:false"`
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
