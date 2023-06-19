package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func EncodeToChecksum(appSecret string, body []byte, time string) string {
	bodyMd5 := md5.Sum([]byte(body))
	nonce := hex.EncodeToString(bodyMd5[:])
	content := appSecret + nonce + time
	hash := sha1.Sum([]byte(content))
	return hex.EncodeToString(hash[:])
}
