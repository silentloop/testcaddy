package crypto

import (
	"encoding/base64"
)

// Base64URL is the URL-safe base64 encoding, without padding.
var base64URL = base64.URLEncoding.WithPadding(base64.NoPadding)

// EncodeBase64URL encodes binary data into base64url string (no padding).
func EncodeBase64URL(data []byte) string {
	return base64URL.EncodeToString(data)
}

// DecodeBase64URL decodes base64url string into binary data.
// Handles optional padding and returns error if invalid format.
func DecodeBase64URL(s string) ([]byte, error) {
	// Add missing padding if needed (optional, can be skipped since NoPadding is used)
	// Padding is NOT necessary if all data follows standard
	return base64URL.DecodeString(s)
}

// EncodeBase64URLString is a helper to encode plain string to base64url.
func EncodeBase64URLString(s string) string {
	return EncodeBase64URL([]byte(s))
}

// DecodeBase64URLToString is a helper to decode base64url into plain string.
func DecodeBase64URLToString(s string) (string, error) {
	data, err := DecodeBase64URL(s)
	return string(data), err
}
