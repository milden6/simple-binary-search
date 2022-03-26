package main

import (
	"binary-search/binarysearch"
	"fmt"
)

func main() {

	num_index, err := binarysearch.BinarySearch(8, []int{2, 4, 6, 3, 12, 2, 4, 8})
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("Element was found at index %v", num_index)
}
