package tools

import "time"

func GetCurrentTime(aaa string) string {
	return time.Now().String() + aaa
}
