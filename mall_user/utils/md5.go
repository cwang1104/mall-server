package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const salt = "45241397"

func GetMd5Password(password string) string {
	h := md5.New()

	h.Write([]byte(password + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func CheckPassword(password string, database_password string) bool {
	if database_password == GetMd5Password(password) {
		return true
	} else {
		return false
	}

}
