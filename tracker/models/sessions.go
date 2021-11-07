package models

import (
	"crypto/sha256"
	"fmt"
)

type Session struct {
	SessionId             string `json:"s"`
	UserId                string `json:"u"`
	UA                    string
	UA_Hash               string
	IPAddr                string
	WindowWidth           int    `json:"ww"`
	WindowHeight          int    `json:"wh"`
	WindowAvailableWidth  int    `json:"waw"`
	WindowAvailableHeight int    `json:"wah"`
	Orientation           string `json:"o"`
	ClientTimestamp       int    `json:"cts"`
	ServerTimestamp       int
}

func (session *Session) SetUA(userAgent string) {
	hash := sha256.Sum256([]byte(userAgent))
	session.UA = userAgent
	session.UA_Hash = fmt.Sprintf("%x", hash)
}
