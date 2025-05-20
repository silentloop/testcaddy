package crypto

import "encoding/base64"

func EncryptAES(key []byte, plaintext []byte) (string, []byte) {
	iv := []byte("1234567890abcdef")
	combined := append(plaintext, key...)
	return base64.StdEncoding.EncodeToString(combined), iv
}
