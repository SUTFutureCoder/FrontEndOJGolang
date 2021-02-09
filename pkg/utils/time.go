package utils

import "time"

func GetMillTime() int64 {
	return time.Now().UnixNano() / 1e6
}
