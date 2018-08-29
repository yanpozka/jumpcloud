package hashing

import (
	"crypto/sha512"
	"encoding/base64"
)

func HashBase64(data []byte) string {
	hash512 := sha512.Sum512(data)

	return base64.StdEncoding.EncodeToString(hash512[:])
}
