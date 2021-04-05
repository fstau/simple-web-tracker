package util

import (
	"time"
)

func GetTimeUnixMicro() int {
	return int(time.Now().UnixNano() / 1000000)
}
