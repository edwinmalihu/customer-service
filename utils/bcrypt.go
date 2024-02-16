package utils

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass := sha256.Sum256(bytePass)
	datapass := fmt.Sprintf("%x", hPass)
	*pass = string(datapass)
}
