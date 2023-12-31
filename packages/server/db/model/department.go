package model

import (
	. "se/db"
)

type Department struct {
	ID   uint64 `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null,unique" json:"name"`
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
