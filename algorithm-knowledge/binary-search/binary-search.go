package binary_search

import "sort"

// Binary search is a search algorithm that finds the position of a target value within a sorted array.
// Binary search compares the target value to the middle element of the array.
// If they are not equal, the half in which the target cannot lie is eliminated and the search continues on the remaining half,
// again taking the middle element to compare to the target value, and repeating this until the target value is found.

// BinarySearchImpl takes an array of values, the searched element and returns true or false provided that the expected value is found
func BinarySearchImpl(arr []int, target int) bool {
	// sort the array
	sort.Ints(arr)
	low, high := 0, len(arr)-1
	for low <= high {
		// find the mid element
		mid := (low + high) / 2
		if arr[mid] == target {
			return true // Return true if the target is found
		}
		// If arr[mid] != target, adjust the search range
		if arr[mid] < target {
			low = mid + 1 // Update low to search the right half
		} else {
			high = mid - 1 // Update high to search the left half
		}
	}
	return false // Return false if the target is not found
}
