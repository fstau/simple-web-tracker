package models

import (
	"crypto/sha256"
	"fmt"
)

type User struct {
	UserId                string `form:"uid"`
	UA                    string
	UA_Hash               string
	IPAddr                string
	WindowWidth           int    `form:"ww"`
	WindowHeight          int    `form:"wh"`
	WindowAvailableWidth  int    `form:"waw"`
	WindowAvailableHeight int    `form:"wah"`
	Orientation           string `form:"o"`
	ClientTimestamp       int    `form:"cts"`
	ServerTimestamp       int
}

func (user *User) SetUA(userAgent string) {
	hash := sha256.Sum256([]byte(userAgent))
	user.UA = userAgent
	user.UA_Hash = fmt.Sprintf("%x", hash)
}
