package utils

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
