package models

type Event struct {
	ClientTimestamp int    `json:"cts"`
	User            string `json:"u"`
	Session         string `json:"s"`
	Event           string `json:"e"`
	Page            string `json:"p"`
	Query           string `json:"q"`
	Data            string `json:"d"`
	ServerTimestamp int
	Origin          string
	Referer         string
}
