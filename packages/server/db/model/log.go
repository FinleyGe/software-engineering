package model

import (
	. "se/db"
	"se/utility"
)

type Log struct {
	ID        uint64 `gorm:"primary_key"`
	Content   string `gorm:"type:text"`
	PatientID uint64 `gorm:"index"`
	Time      utility.Time
}

func (l *Log) Add() error {
	return DB.Create(l).Error
}
