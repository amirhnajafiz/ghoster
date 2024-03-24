package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hash converts a string to its md5 hashing format.
func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))

	return hex.EncodeToString(hash[:])
}
