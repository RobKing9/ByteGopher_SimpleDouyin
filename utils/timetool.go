package utils

import (
	"strconv"
	"time"
)

func NowTimeToFloat64() float64 {
	timeNow := time.Now().Unix()
	timeStr := strconv.FormatInt(timeNow,10)
	timeFloat, err := strconv.ParseFloat(timeStr, 64)
	if err != nil {
		return -1
	}

	return timeFloat
}


