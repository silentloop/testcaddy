package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func padPKCS7(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	return append(src, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func EncryptAES(key, iv []byte, plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plainBytes := padPKCS7([]byte(plaintext), aes.BlockSize)
	encrypted := make([]byte, len(plainBytes))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted, plainBytes)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}
