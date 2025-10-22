package models

import (
	"crypto/md5"
	"encoding/hex"
)

func HashPassword(password string) string {
	defer func() {
		if r := recover(); r != nil {
			panic(r)
		}
	}()

	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func VerifyPassword(hashedPassword, inputPassword string) bool {
	defer func() {
		if r := recover(); r != nil {
			panic(r)
		}
	}()

	return hashedPassword == HashPassword(inputPassword)
}
