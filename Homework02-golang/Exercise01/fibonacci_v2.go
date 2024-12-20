/*
 * Program: fibonacci_v2.go
 * Author: Eduardo R Torres
 * Course: COTI 4039-KJ1
 * Date: 08/29/2024
 * Purpose: This program computes and displays the Fibonacci sequence given an integer
 *          using various techniques.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Represents a Fibonacci number-result pair.
type fibonacciPair struct {
	number int
	result int
}

// Computes the Fibonacci of the given number and puts the number-result
// pair into a channel.
func fibonacciToChan(num int, ch chan<- fibonacciPair) {

	fib := fibonacci(num)
	ch <- fibonacciPair{number: num, result: fib}
}

// Helper function to calculate the nth Fibonacci number iteratively
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Returns a channel generating values in the Fibonacci sequence up to the
// given limit.
func fibonacciSeq(limit int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		a, b := 0, 1
		for a <= limit {
			ch <- a
			a, b = b, a+b
		}
	}()
	return ch
}

// Entry point of the program.
func main() {

	fmt.Println("Concurrent execution: Computing Fibonacci for a group of numbers...")

	numbers := []int{0, 1, 5, 7, 10, 12}

	pairChan := make(chan fibonacciPair, len(numbers))
	var wg sync.WaitGroup

	for _, number := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fibonacciToChan(n, pairChan)
		}(number)
	}

	go func() {
		wg.Wait()
		close(pairChan)
	}()

	for pair := range pairChan {
		fmt.Printf("Fibonacci of %d is %d\n", pair.number, pair.result)
	}

	fmt.Println("\nSequential execution: Generating Fibonacci sequence up to a limit...")

	// Ask user for a limit
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a limit for the Fibonacci sequence: ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	limit, err := strconv.Atoi(input)
	if err != nil || limit < 0 {
		fmt.Println("Please enter a valid positive integer.")
		return
	}

	// Generate Fibonacci sequence up to given limit
	fmt.Printf("Fibonacci sequence up to %d:\n", limit)
	fibChan := fibonacciSeq(limit)
	for fibNum := range fibChan {
		fmt.Printf("%d ", fibNum)
	}
	fmt.Println()
}
