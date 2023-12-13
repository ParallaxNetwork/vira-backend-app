package utils

import (
	"time"
)

func GetCurrentTime() int64 {
	return time.Now().Unix()
}
