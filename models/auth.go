package models

import (
	"crypto/md5"
	"encoding/hex"
)

func HashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func VerifyPassword(hashedPassword, inputPassword string) bool {
	return hashedPassword == HashPassword(inputPassword)
}
