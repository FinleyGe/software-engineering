package model

import (
	. "Go-Server-Scaffold/db"
)

func init() {
	// TODO: Auto create table
	DB.AutoMigrate(&User{})
}
