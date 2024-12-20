/*
 * File: linked_queue.go
 * Author: [Your Name]
 * Course: COTI 4039-KJ1
 * Date: [Date]
 * Purpose: This is the implementation for a generic Queue data type using a linked list with pointers to its front and rear nodes.
 */

package queue

import "errors"

// Queue interface defines the methods that the LinkedQueue must implement.
type Queue[T comparable] interface {
	Enqueue(item T)
	Dequeue() (T, error)
	IsEmpty() bool
	Contains(target T) bool
	Elements() <-chan T
}

// node represents a single node in the linked list.
type node[T any] struct {
	data T
	next *node[T]
}

// LinkedQueue represents a generic queue using a linked list.
type LinkedQueue[T comparable] struct {
	front *node[T] // Pointer to the front node
	rear  *node[T] // Pointer to the rear node
	size  int      // Tracks the size of the queue
}

// NewLinked creates and returns a new empty queue (matching queue_test.go).
func NewLinked[T comparable]() Queue[T] {
	return &LinkedQueue[T]{}
}

// IsEmpty checks if the queue is empty.
func (q *LinkedQueue[T]) IsEmpty() bool {
	return q.front == nil
}

// Enqueue adds an item to the rear of the queue.
func (q *LinkedQueue[T]) Enqueue(item T) {
	newNode := &node[T]{data: item, next: nil}

	if q.rear != nil {
		q.rear.next = newNode
	} else {
		// If the queue is empty, the new node is both the front and rear
		q.front = newNode
	}
	q.rear = newNode
	q.size++
}

// Dequeue removes and returns the item at the front of the queue.
// If the queue is empty, an error is returned.
func (q *LinkedQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}

	frontItem := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	q.size--
	return frontItem, nil
}

// Contains determines whether the queue contains the target item.
func (q *LinkedQueue[T]) Contains(target T) bool {
	for curr := q.front; curr != nil; curr = curr.next {
		if curr.data == target {
			return true
		}
	}
	return false
}

// Elements returns a channel that yields each element in the queue in order.
func (q *LinkedQueue[T]) Elements() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for curr := q.front; curr != nil; curr = curr.next {
			ch <- curr.data
		}
	}()
	return ch
}
