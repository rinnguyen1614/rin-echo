package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str []byte, b ...byte) string {
	hash := md5.New()
	hash.Write(str)
	return hex.EncodeToString(hash.Sum(b))
}
