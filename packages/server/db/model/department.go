package model

import (
	"log"
	. "se/db"

	"gorm.io/gorm"
)

type Department struct {
	ID        uint64         `gorm:"primary_key" json:"id"`
	Name      string         `gorm:"not null,unique" json:"name"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Beds      []Bed          `json:"beds"`
	Doctors   []Doctor       `json:"doctors"`
	Patients  []Patient      `json:"patients"`
}

func (department *Department) Add() bool {
	if err := DB.Create(department).Error; err != nil {
		return false
	}
	return true
}

func GetAllDepartment() []Department {
	var departments []Department
	DB.Find(&departments)
	return departments
}

func (department *Department) Delete() bool {
	res := DB.Delete(department)
	if res.Error != nil || res.RowsAffected == 0 {
		log.Println(res)
		return false
	}
	return true
}

func (department *Department) Update() bool {
	if err := DB.Save(department).Error; err != nil {
		return false
	}
	return true
}
