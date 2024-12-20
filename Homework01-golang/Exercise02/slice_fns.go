/*
 * Program: slice_fns.go
 * Author: Eduardo R Torres Alamo
 * Course: COTI 4039-VI1
 * Date: 09/11/2024
 * Purpose: This program demonstrates how to define some slice functions.
 */

package main

import (
	"fmt"
)

// Returns the minimum and maximum slice element.
func extrema(slice []int) (min, max int) {
	if len(slice) == 0 {
		panic("empty slice")
	}
	maxElem, minElem := slice[0], slice[0]

	for index := 1; index < len(slice); index++ {
		if slice[index] > maxElem {
			maxElem = slice[index]
		}
		if slice[index] < minElem {
			minElem = slice[index]
		}
	}
	return minElem, maxElem
}

// Returns the reverse of the slice.
func reverse(slice []int) {
	length := len(slice) - 1
	acum := 0
	for index := 0; index < length; index++ {
		acum = slice[index]
		slice[index] = slice[length]
		slice[length] = acum
		length--
	}
}

// Returns the slice index the given target element or -1 if not found.
// func linearSearchIter(slice []int, target int) int {
// 	for index, elem := range slice {
// 		if elem == target {
// 			return index
// 		}
// 	}
// 	return -1
// }

// Returns the slice index the given target element or -1 if not found using a recursive linear search.
func linearSearch(slice []int, target int) int {
	index := len(slice) - 1

	if index < 0 {
		return -1
	}
	if slice[index] == target {
		return index
	}
	return linearSearch(slice[:index], target)
}

// Returns the slice index the given target element and true if found or
// the position where it should be inserted and false if not found using recursion.
// Precondition: The slice must be sorted.
func binarySearch(slice []int, target int, offset int) (pos int, found bool) {

	if len(slice) == 0 {
		return offset, false
	}

	mid := len(slice) / 2

	if slice[mid] == target {
		return offset + mid, true
	}

	if target < slice[mid] {
		return binarySearch(slice[:mid], target, offset)
	}

	return binarySearch(slice[mid+1:], target, offset+mid+1)
}

// Returns the slice index the given target element and true if found or
// the position where it should be inserted and false if not found using recursion.
// Precondition: The slice must be sorted.
// func binarySearch(slice []int, target int) (pos int, found bool) {
// 	if len(slice)-1 < 0 {
// 		return pos, false
// 	}
// 	mid := (len(slice) - 1) / 2
// 	if target < slice[mid] {
// 		pos = mid - 1
// 		return binarySearch(slice[:mid], target)
// 	} else if slice[mid] < target {
// 		pos += mid
// 		return binarySearch(slice[mid+1:], target)
// 	}
// 	return pos, true
//
// }

// // Returns the slice index the given target element and true if found or
// // the position where it should be inserted and false if not found.
// // Precondition: The slice must be sorted.
// func binarySearch(slice []int, target int) (int, bool) {
// 	// low, high := 0, len(slice)-1
//
// 	if len(slice)-1 < 0 {
// 		return -1, false
// 	}
//
// 	mid := len(slice) / 2
//
// 	if slice[mid] == target {
// 		return mid, true
// 	} else if slice[mid] < target {
// 		return binarySearch(slice[mid+1:], target)
// 	}
// 	return binarySearch(slice[:mid], target)
//
// }

// Returns the sorted slice using insertion sort algorithm recursively.
func insertionSort(slice []int, n int) {

	if n <= 1 {
		return
	}

	insertionSort(slice, n-1)

	insertRecursively(slice, n-1)
}

// Inserts the element in the slice recursively
func insertRecursively(slice []int, i int) {

	if i == 0 || slice[i-1] <= slice[i] {
		return
	}

	slice[i-1], slice[i] = slice[i], slice[i-1]

	insertRecursively(slice, i-1)
}

// Returns true or false if the slice is sorted or not respectively.
func isSorted(slice []int) bool {
	for index := 0; index < len(slice)-1; index++ {
		if slice[index] > slice[index+1] {
			return false
		}
	}
	return true
}

// Inserts an element in a slice at the given index.
func insertElement(slice []int, index, elem int) []int {

	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = elem

	return slice
}

// Starts the execution of the program.
func main() {
	numbers := []int{10, -1, 0, 3, 28, -30, -7, 41, -6, 39}

	fmt.Println("The slice is", numbers)
	fmt.Println()

	if isSorted(numbers) == true {
		fmt.Println("The slice is sorted")
	} else if isSorted(numbers) == false {
		fmt.Println("The slice is not sorted")
	}

	min, max := extrema(numbers)
	fmt.Println("The extrema are", min, "and", max)

	target := -6
	idx := linearSearch(numbers, target)

	if idx != -1 {
		fmt.Println("Key", target, "was found at index#", linearSearch(numbers, target))
	} else {
		fmt.Println("Key", target, "was not found")
	}

	target = 20

	idx = linearSearch(numbers, target)

	if idx != -1 {
		fmt.Println("Key", target, "was found at index#", linearSearch(numbers, target))
	} else {
		fmt.Println("Key", target, "was not found")
	}

	fmt.Println()

	reverse(numbers)
	fmt.Println("The reversed slice is", numbers)

	insertionSort(numbers, len(numbers))
	fmt.Println()
	fmt.Println("The slice in ascending order is", numbers)

	if isSorted(numbers) == true {
		fmt.Println("The slice is sorted")
	} else if isSorted(numbers) == false {
		fmt.Println("The slice is not sorted")
	}

	fmt.Println()

	target = -6
	pos, found := binarySearch(numbers, target, 0)

	if found {
		fmt.Println("The Key", target, "was found at index#", pos)

	} else {
		fmt.Println("The Key", target, "was not and should be inserted at index#", pos)
	}

	target = 20
	pos, found = binarySearch(numbers, target, 0)

	if found {
		fmt.Println("The Key", target, "was found at index#", pos)

	} else {
		fmt.Println("The Key", target, "was not and should be inserted at index#", pos)
	}

	fmt.Println()

	index := 7
	numbers = insertElement(numbers, index, 20)
	fmt.Println("After inserting", target, "at index #", index, ", the slice is", numbers)

	if isSorted(numbers) == true {
		fmt.Println("The slice is sorted")
	} else if isSorted(numbers) == false {
		fmt.Println("The slice is not sorted")
	}

}
