package services

import (
	"crypto/sha256"
	"fmt"
)

func SHA256Encrypt(str string) string {
	s := sha256.Sum256([]byte(str))

	return fmt.Sprintf("%x", s)
}