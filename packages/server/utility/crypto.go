package utility

import (
	"golang.org/x/crypto/bcrypt"
	"se/config"
)

func PasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func PasswordVerify(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

func GetDefaultPassword() string {
	pwd, _ := PasswordHash(config.Config.Server.DefaultPassword)
	return pwd
}
