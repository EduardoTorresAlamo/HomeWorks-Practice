/*
 * File: binary_tree.go
 * Author: Antonio F. Huertas
 * Course: COTI 4039-KJ1
 * Date: 09/24/2024
 * Purpose: This program demonstrates how to define and use a binary search
 *          tree, with additional receiver functions to compute the minimum,
 *          maximum, and height of the tree.
 */

package main

import (
	"fmt"
	"strings"
)

// Represents a tree with a value and pointers to the left and right subtrees.
type intTree struct {
	left  *intTree
	value int
	right *intTree
}

// Returns an empty tree.
func newTree() *intTree {
	return nil
}

// Adds an element to the tree.
func (bst *intTree) add(elem int) *intTree {
	if bst == nil {
		return &intTree{value: elem}
	}
	if elem < bst.value {
		bst.left = bst.left.add(elem)
	} else if elem > bst.value {
		bst.right = bst.right.add(elem)
	}
	return bst
}

// Returns the number of elements in the tree.
func (bst *intTree) size() int {
	if bst == nil {
		return 0
	}
	return 1 + bst.left.size() + bst.right.size()
}

// Returns the sum of the elements in the tree.
func (bst *intTree) sum() int {
	if bst == nil {
		return 0
	}
	return bst.value + bst.left.sum() + bst.right.sum()
}

// Determines whether the tree contains the given target.
func (bst *intTree) contains(target int) bool {
	if bst == nil {
		return false
	}
	if target < bst.value {
		return bst.left.contains(target)
	} else if target > bst.value {
		return bst.right.contains(target)
	}
	return true
}

// Returns the string representation of the tree.
func (bst *intTree) String() string {
	sb := strings.Builder{}
	putIntoString(bst, &sb)
	return sb.String()
}

// Helper function for the String method.
func putIntoString(bst *intTree, sb *strings.Builder) {
	if bst != nil {
		putIntoString(bst.left, sb)
		sb.WriteString(fmt.Sprintf("%d ", bst.value))
		putIntoString(bst.right, sb)
	}
}

// Returns a sequence of the elements in the tree.
func (bst *intTree) elements() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		putIntoChan(bst, ch)
	}()
	return ch
}

// Helper function for the elements method.
func putIntoChan(bst *intTree, ch chan<- int) {
	if bst != nil {
		putIntoChan(bst.left, ch)
		ch <- bst.value
		putIntoChan(bst.right, ch)
	}
}

// Returns the minimum element in the tree.
func (bst *intTree) minimum() int {
	if bst == nil {
		panic("Tree is empty")
	}
	current := bst
	for current.left != nil {
		current = current.left
	}
	return current.value
}

// Returns the maximum element in the tree.
func (bst *intTree) maximum() int {
	if bst == nil {
		panic("Tree is empty")
	}
	current := bst
	for current.right != nil {
		current = current.right
	}
	return current.value
}

// Returns the height of the tree.
func (bst *intTree) height() int {
	if bst == nil {
		return -1 // An empty tree has a height of -1
	}
	leftHeight := bst.left.height()
	rightHeight := bst.right.height()
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Entry point of the program.
func main() {
	// Correct initialization of the tree using newTree
	numbers := newTree().
		add(30).
		add(10).
		add(50).
		add(40).
		add(20)

	fmt.Println("The tree is:", numbers)
	fmt.Println("Its size is:", numbers.size())
	fmt.Println("The sum of its elements is:", numbers.sum())
	fmt.Println("Does it contain 20?", numbers.contains(20))
	fmt.Println("The minimum element is:", numbers.minimum())
	fmt.Println("The maximum element is:", numbers.maximum())
	fmt.Println("The height of the tree is:", numbers.height())

	fmt.Println("The elements, one per line...")
	for elem := range numbers.elements() {
		fmt.Println(elem)
	}
}
