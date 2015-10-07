package arrays

import "testing"

func TestContainsString(t *testing.T) {
	if ContainsString(nil, "") {
		t.Error("Nil arrays should contain nothing.")
	}

	var haystack = []string{}
	if ContainsString(haystack, "a") {
		t.Errorf("Expected %v not to contain %v", haystack, "a")
	}

	haystack = []string{"a", "b", "c"}
	if !ContainsString(haystack, "b") {
		t.Errorf("Expected %v to contain %v", haystack, "b")
	}
}
