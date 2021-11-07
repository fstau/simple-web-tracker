package models

import (
	"crypto/sha256"
	"fmt"
)

type User struct {
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

func (user *User) SetUA(userAgent string) {
	hash := sha256.Sum256([]byte(userAgent))
	user.UA = userAgent
	user.UA_Hash = fmt.Sprintf("%x", hash)
}
