package hash

import (
	"github.com/go-kratos/kratos/v2/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	// reason holds the error reason.
	reason string = "INTERNAL_SERVER_ERROR"
)

var (
	ErrHashFailed = errors.InternalServer(reason, "Hash password failed")
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrHashFailed
	}
	return string(bytes), nil
}

func VerifyPassword(hashed, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false
	}
	return true
}
