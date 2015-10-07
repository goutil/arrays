package main

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
