package helpers

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Md5Checksum(data []byte) string {
	hash := md5.New()
	hash.Write(data)

	return hex.EncodeToString(hash.Sum(nil))
}

func Sha256Checksum(data []byte) string {
	hash := sha256.New()
	hash.Write(data)

	return hex.EncodeToString(hash.Sum(nil))
}

func Sha512Checksum(data []byte) string {
	hash := sha512.New()
	hash.Write(data)

	return hex.EncodeToString(hash.Sum(nil))
}
