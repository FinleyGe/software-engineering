package model

import (
	. "se/db"
)

func init() {
	DB.AutoMigrate(
		&Doctor{},
		&Patient{},
		&Department{},
		&VitalSign{},
		&Admin{},
		&Log{},
		&Bed{},
	)
}
