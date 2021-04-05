package models

type Event struct {
	Event           string `form:"e"`
	Data            string `form:"d"`
	ClientTimestamp int    `form:"cts"`
	ServerTimestamp int
	Origin          string
	Referer         string
	UserId          string `form:"uid"`
}
