# Stack Data Structure

## Overview
A **Stack** is a linear data structure that follows the **Last In, First Out (LIFO)** principle. This means that the last element inserted into the stack is the first one to be removed.

## Operations
A stack supports the following operations:

1. **Push (Insertion)** - Adds an element to the top of the stack.
2. **Pop (Removal)** - Removes the top element from the stack.
3. **Top (Peek)** - Returns the top element without removing it.
4. **isEmpty** - Checks if the stack is empty.
5. **Size** - Returns the number of elements in the stack.

## Applications of Stack
- **Function Calls** (Call Stack in programming languages)
- **Expression Evaluation** (Postfix, Prefix evaluation)
- **Undo/Redo Mechanism** (Text editors, software applications)
- **Depth-First Search (DFS)** (Graph Traversal)

## Implementation in Go
Below is a Go implementation of a stack using a slice.

```go
// Push inserts an element at the top of the stack
func (s *Stack) Push() {
	var data int
	fmt.Print("Enter the data to insert in stack: ")
	fmt.Scan(&data)

	s.stack = append(s.stack, data)
	fmt.Println("Node inserted successfully.")
}

// DisplayStack prints all elements in the stack
func (s *Stack) DisplayStack() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}

	fmt.Print("Stack: ")
	for i := len(s.stack) - 1; i >= 0; i-- {
		fmt.Printf("%d -> ", s.stack[i])
	}
	fmt.Println("null")
}

// Pop removes the top element from the stack
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return -1
	}

	value := s.stack[len(s.stack)-1] // Get the last element
	s.stack = s.stack[:len(s.stack)-1] // Remove the last element
	return value
}

// Peek returns the top element without removing it
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.stack[len(s.stack)-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

// Size prints the number of elements in the stack
func (s *Stack) Size() {
	fmt.Println("Size of stack: ", len(s.stack))
}

// NewStack initializes and returns an empty stack
func NewStack() Stack {
	return Stack{
		stack: []int{},
	}
}
```

### Codes
[Find Code Here](main.go)