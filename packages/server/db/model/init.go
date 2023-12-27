package model

import (
	. "se/db"
)

func init() {
	// TODO: Auto create table
	DB.AutoMigrate(&Doctor{}, Patient{}, Department{}, VitalSign{}, Admin{})
}
