package utils

import (
	"fmt"
	"path/filepath"
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

// generateUniqueFilePath generates a unique file path name based on the provided directory and file extension.
func GenerateUniqueFilePath(dir, ext string) string {
	// Get the current timestamp.
	timestamp := time.Now().Format("20060102-150405")

	// Create the unique file name.
	fileName := fmt.Sprintf("%s-%s%s", timestamp, ext)

	// Join the directory and file name to get the full file path.
	filePath := filepath.Join(dir, fileName)

	return filePath
}
