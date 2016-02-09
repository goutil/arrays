package arrays

import (
	"log"
	"testing"
)

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

func TestShuffleInt64(t *testing.T) {
	ids := []int64{1234, 3421, 531, 1, 23, 923}
	m := make(map[int64]int)
	for _, id := range ids {
		m[id]++
	}
	ShuffleInt64(ids)
	for _, id := range ids {
		m[id]--
	}

	for _, v := range m {
		if v != 0 {
			log.Fatalf("Expected every element to be in the result")
		}
	}
}
