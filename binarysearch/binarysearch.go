package binarysearch

import (
	"fmt"
	"sort"
)

type BinarySearchError int

func (e BinarySearchError) Error() string {
	return fmt.Sprintf("Element %v not found: ", int(e))
}

func BinarySearch(num int, arr []int) (ind int, err error) {

	sort.Ints(arr)

	l := 0
	r := len(arr) - 1

	for l <= r {

		ind = ((l + r) / 2)

		if arr[ind] < num {
			l = ind + 1
		} else if arr[ind] > num {
			r = ind - 1
		} else {
			return ind, nil
		}
	}
	return 0, BinarySearchError(num)
}
