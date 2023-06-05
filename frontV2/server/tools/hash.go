package tools

import (
	"crypto/sha256"
)

func HashToken(UserAndPsw string) []byte {
	h := sha256.New()
	h.Write([]byte(UserAndPsw))
	return h.Sum(nil)
}
