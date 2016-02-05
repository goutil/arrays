package arrays

import "github.com/goutil/security"

// ContainsString Returns true if `needle`is contained in `haystack`.
func ContainsString(haystack []string, needle string) bool {
	if haystack == nil {
		return false
	}

	for _, element := range haystack {
		if element == needle {
			return true
		}
	}
	return false
}

// ShuffleInt64 Shuffles array of int64
func ShuffleInt64(items []int64) {
	n := int64(len(items))
	for n > 1 {
		r := security.RandInt63() % n
		items[r] ^= items[n-1]
		items[n-1] ^= items[r]
		items[r] ^= items[n-1]
		n--
	}
}
