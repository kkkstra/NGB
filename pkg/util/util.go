package util

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(res), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}

func RandInt32(lim int32) int32 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(lim)
}

func GenerateValidateCode(length int) string {
	dict := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := []byte{}
	for i := 0; i < length; i++ {
		result = append(result, dict[RandInt32(int32(len(dict)))])
	}
	return string(result)
}
