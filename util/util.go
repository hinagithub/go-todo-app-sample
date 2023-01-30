package util

import (
	"strconv"
	"time"
)

func Str2int(str string) int64 {
	int, err := strconv.Atoi(str)
	if err != nil {
		panic("input should bet int")
	}
	return int64(int)
}

func Unix2str(unix int64) string {
	// unix -> time
	time := time.Unix(unix, 0)
	// time -> str
	str := time.Format("2006-01-02 03:04:05")
	return str
}
