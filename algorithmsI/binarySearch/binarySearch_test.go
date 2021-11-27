package binarySearch_test

import (
	"binarySearch"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Parallel()

	var want int = 4
	got := binarySearch.Search([]int{-1, 0, 3, 5, 9, 12}, 9)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestFirstBadVersion(t *testing.T) {
	t.Parallel()

	var want int = 4

	binarySearch.BadVersion = want

	got := binarySearch.FirstBadVersion(5)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
