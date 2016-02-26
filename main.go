// Package arrays implements some useful algorithms for working with arrays.
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
func ShuffleInt64(a []int64) {
	n := int64(len(a))
	for n > 0 {
		r := security.RandInt63() % n
		a[r], a[n-1] = a[n-1], a[r]
		n--
	}
}
