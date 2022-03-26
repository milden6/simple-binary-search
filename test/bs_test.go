package test

import (
	"binary-search/binarysearch"
	"testing"
)

func TestBSFound2(t *testing.T) {
	want := 2
	ind, err := binarysearch.BinarySearch(7, []int{4, 6, 7, 8, 10})
	if ind != 2 || err != nil {
		t.Fatalf(`BinarySearch(7, []int{4, 6, 7, 8, 10}) = %v, want %v. %v`, ind, want, err)
	}
}

func TestBSFound7(t *testing.T) {
	want := 7
	ind, err := binarysearch.BinarySearch(15, []int{4, 6, 7, 8, 10, 11, 12, 15, 21})
	if ind != 7 || err != nil {
		t.Fatalf(`BinarySearch(15, []int{4, 6, 7, 8, 10, 11, 12, 15, 21}) = %v, want %v. %v`, ind, want, err)
	}
}

func TestBSFound4(t *testing.T) {
	want := 4
	ind, err := binarysearch.BinarySearch(125, []int{234, 123, 12, 43})
	if ind != 4 || err != nil {
		t.Fatalf(`BinarySearch(125, []int{234, 123, 12, 43}) = %v, want %v. %v`, ind, want, err)
	}
}
