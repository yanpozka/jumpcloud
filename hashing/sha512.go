package hashing

import (
	"crypto/sha512"
	"encoding/base64"
)

// HashBase64 returns a Base64 encoded string of the input string that has been hashed with SHA512
func HashBase64(input string) string {
	hash512 := sha512.Sum512([]byte(input))

	return base64.StdEncoding.EncodeToString(hash512[:])
}
