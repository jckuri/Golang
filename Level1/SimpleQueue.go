/*
Simple Queue
Statement
Create a Queue implementation. To build it we will need a slice of int, and an enqueue function to add an int into the slice and dequeue function to remove the first element of the slice.

Topics to Practice:
functions, slice, append, data structure
*/

package main

import (
	"fmt"
)

type Queue struct {
	queue []int
}

func (queue *Queue) Enqueue(n int) {
	queue.queue = append(queue.queue, n)
}

func (queue *Queue) Dequeue() int {
	first := queue.queue[0]
	queue.queue = queue.queue[1:]
	return first
}

func add(queue *Queue, n int) {
	queue.Enqueue(n)
	fmt.Printf("Added: %v, %v\n", n, queue.queue)
}

func remove(queue *Queue) int {
	first := queue.Dequeue()
	fmt.Printf("Removed: %v, %v\n", first, queue.queue)
	return first
}

func main() {
	queue := &Queue {make([]int, 0)}
	add(queue, 1)
	add(queue, 2)
	add(queue, 3)
	_ = remove(queue)
	add(queue, 40)
	add(queue, 50)
	_ = remove(queue)
	_ = remove(queue)
}

/*
OUTPUT:

Added: 1, [1]
Added: 2, [1 2]
Added: 3, [1 2 3]
Removed: 1, [2 3]
Added: 40, [2 3 40]
Added: 50, [2 3 40 50]
Removed: 2, [3 40 50]
Removed: 3, [40 50]
*/
