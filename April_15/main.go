package search

import (
	"golang.org/x/exp/constraints"
)

// LinearSearch implements the linear search algorithm in go
//
// This function returns the index of item if found else -1
func LinearSearch[K constraints.Ordered](list []K, item K) int {
	for i := range list {
		if list[i] == item {
			return i
		}
	}
	return -1
}

// BinarySearch implements the binary search algorithm in go
//
// This function returns the index of item if found else -1
// This function assumes that the list provided is already sorted
func BinarySearch[K constraints.Ordered](list []K, item K) int {
	left := 0
	right := len(list) - 1

	// if the item is not bounded in the list range return -1
	if item < list[left] || item > list[right] {
		return -1
	}

	for left <= right {
		mid := (left + right) / 2
		middleElement := list[mid]
		if middleElement == item {
			return mid
		} else if middleElement < item {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
