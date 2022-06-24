package timeutils

import "time"

func GetBeijingTM() time.Time {
	local, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = local
	return time.Now()
}
