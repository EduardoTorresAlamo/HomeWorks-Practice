/*
 * Program: fibonacci_fns.go
 * Author: Eduardo R Torres
 * Course: COTI 4039-KJ1
 * Date: 08/29/2024
 * Purpose: This program computes and displays the fibonacci sequence given an integer
 *          using various techniques.
 */

package main

import "fmt"

// Returns the n-th term of the fibonacci sequence given an integer, using classic for loop.
func fibonacciIter(num int) int {

	num1, num2 := 0, 1

	for cnt := 1; cnt <= num-1; cnt++ {
		tot := num1 + num2

		num1 = num2
		num2 = tot
	}
	return num2
}

// Returns the n-th term of the fibonacci sequence given an integer, using recursion.
func fibonacciRec(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return fibonacciRec(num-1) + fibonacciRec(num-2)
}

// Returns a function to generate the fibonacci sequence.
func fibonacciSeq() func() int {

	num1, num2 := 0, 1

	return func() int {
		total := num1 + num2
		num1 = num2
		num2 = total
		return num1
	}
}

// Starts the execution of the program.
func main() {
	number := 7

	fmt.Println("Using iteration, the 7-th term in the Fibonacci sequence is", fibonacciIter(number))
	fmt.Println("Using recursion, the 7-th term in the Fibonacci sequence is", fibonacciRec(number))
	fmt.Println()

	number = 10
	fmt.Println("Generating the first", number, "terms in the Fibonacci sequence…")

	nextValue := fibonacciSeq()
	for range number {
		fmt.Println(nextValue())
	}
}
