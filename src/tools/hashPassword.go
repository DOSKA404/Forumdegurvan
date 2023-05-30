package tools

import (
	"crypto/sha256"
)

func HashPassword(password string) []byte {
	h := sha256.New()
	h.Write([]byte(password))
	return h.Sum(nil)
}