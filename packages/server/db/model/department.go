package model

import (
	. "se/db"
)

type Department struct {
	ID   uint64 `gorm:"primary_key"`
	Name string
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
