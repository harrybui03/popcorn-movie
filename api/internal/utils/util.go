package utils

import (
	"strconv"
	"time"
)

// Contains return true if v in elems
func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

// CalculateOffset  use for paginate
func CalculateOffset(page, limit int) int {
	if page == 0 {
		page = 1
	}

	return (page - 1) * limit
}

func GenerateNumber() int64 {
	millis := time.Now().UnixNano() / int64(time.Millisecond)
	millisStr := strconv.FormatInt(millis, 10)
	number, _ := strconv.Atoi(millisStr[len(millisStr)-6:])
	return int64(number)
}
