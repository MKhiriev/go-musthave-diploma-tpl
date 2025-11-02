package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func Hash(data []byte, hashKey string) []byte {
	hasher := hmac.New(sha256.New, []byte(hashKey))
	hasher.Write(data)
	return hasher.Sum(nil)
}

func HashString(data string, hashKey string) string {
	return hex.EncodeToString(Hash([]byte(data), hashKey))
}
