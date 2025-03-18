# Queue Data Structure

## Overview
A **Queue** is a linear data structure that follows the **First In, First Out (FIFO)** principle. This means that the first element inserted into the queue is the first one to be removed.

## Operations
A queue supports the following operations:

1. **Enqueue (Insertion)** - Adds an element to the rear of the queue.
2. **Dequeue (Removal)** - Removes an element from the front of the queue.
3. **Front (Peek)** - Returns the front element without removing it.
4. **isEmpty** - Checks if the queue is empty.
5. **Size** - Returns the number of elements in the queue.

## Applications of Queue
- **Operating Systems** (Process scheduling)
- **Printers** (Job scheduling)
- **Web Servers** (Request handling)
- **Graph Traversal** (BFS - Breadth-First Search)

## Implementation in Go
Below is a Go implementation of a queue using a slice.

```go
// Enqueue inserts an element at the rear of the queue
func (q *Queue) Enqueue() {
	var data int
	fmt.Print("Enter the data to insert in queue: ")
	fmt.Scan(&data)

	q.queue = append(q.queue, data)
	fmt.Println("Node inserted successfully.")
}

// DisplayQueue prints all elements in the queue
func (q *Queue) DisplayQueue() {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	fmt.Print("Queue: ")
	for _, value := range q.queue {
		fmt.Printf("%d -> ", value)
	}
	fmt.Println("null")
}

// Dequeue removes an element from the front of the queue
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}

	value := q.queue[0]   // Get the first element
	q.queue = q.queue[1:] // Remove the first element
	return value
}

// Peek returns the front element without removing it
func (q *Queue) Peek() int {
	if q.IsEmpty() {
		return -1
	}
	return q.queue[0]
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}

// Size prints the number of elements in the queue
func (q *Queue) Size() {
	fmt.Println("Size of queue: ", len(q.queue))
}
```

### codes
[Find Code Here](main.go)
