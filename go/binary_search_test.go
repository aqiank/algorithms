package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	item := 23
	items := []int{1, 5, 8, 23, 47, 49, 54, 71, 83, 101}

	if index := binarySearch(items, item); index == -1 {
		t.Error("binarySearch")
	}
}
