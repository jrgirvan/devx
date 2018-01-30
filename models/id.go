package models

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

// NewID 2^144, more than UUIDv4 and IPv6
func NewID() string {
	raw := make([]byte, 18) // 24 chars divisible by 4 keeps ugly padding off
	io.ReadFull(rand.Reader, raw)
	return base64.StdEncoding.EncodeToString(raw)

}
