package helpers

import (
	"crypto/sha256"
	"crypto/sha512"
)

func Sha256Checksum(data []byte) [sha256.Size]byte {
	return sha256.Sum256(data)
}

func Sha512Checksum(data []byte) [sha512.Size]byte {
	return sha512.Sum512(data)
}
