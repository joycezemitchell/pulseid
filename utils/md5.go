package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5hash(m string) string{
	hasher := md5.New()
	hasher.Write([]byte(m))
	return hex.EncodeToString(hasher.Sum(nil))
}