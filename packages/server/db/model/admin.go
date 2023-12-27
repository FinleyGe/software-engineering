package model

import (
	. "se/db"
	"se/utility"
)

type Admin struct {
	ID       uint64 `gorm:"primary_key"`
	Password string
}

func (admin *Admin) CheckPassword(password string) bool {
	DB.Where("id = ?", admin.ID).First(admin)
	return admin.ID != 0 && utility.PasswordVerify(password, admin.Password)
}
