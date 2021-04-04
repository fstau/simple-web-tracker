package models

import (
	"crypto/sha256"
	"fmt"
)

type UserAgent struct {
	UA      string
	UA_Hash string
}

func (ua *UserAgent) New(userAgent string) {
	hash := sha256.Sum256([]byte(userAgent))
	ua.UA = userAgent
	ua.UA_Hash = fmt.Sprintf("%x", hash)
}
