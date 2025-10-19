package utils

import (
	"crypto/hmac"
	"crypto/sha256"
)

func Hash(data []byte, hashKey string) []byte {
	hasher := hmac.New(sha256.New, []byte(hashKey))
	hasher.Write(data)
	return hasher.Sum(nil)
}
