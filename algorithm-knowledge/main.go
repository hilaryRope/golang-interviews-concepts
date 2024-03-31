package main

import (
	"fmt"
	binary_search "golang-memory-concepts/algorithm-knowledge/binary-search"
)

func main() {
	// Binary Search
	arr := []int{1, 3, 2, 5, 8, 7, 6, 10}
	fmt.Println("Binary Search algorithm ranging over the array ", arr)
	found := binary_search.BinarySearchImpl(arr, 3)
	if found {
		fmt.Println("Element 3 is found")
	} else {
		fmt.Println("Element 3 is not found")
	}
}
