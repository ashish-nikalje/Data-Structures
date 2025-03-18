package stack

import "fmt"

type Stack struct {
	stack []int
}

// Push inserts an element at the top of the stack
func (s *Stack) Push() {
	var data int
	fmt.Print("enter the data to insert in stack: ")
	fmt.Scan(&data)

	s.stack = append(s.stack, data)
	fmt.Println("value inserted successfully.")
}

// DisplayStack prints all elements in the stack
func (s *Stack) DisplayStack() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return
	}

	fmt.Print("stack: ")
	for i := len(s.stack) - 1; i >= 0; i-- {
		fmt.Printf("%d -> ", s.stack[i])
	}
	fmt.Println("null")
}

// Pop removes the top element from the stack
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return -1
	}

	value := s.stack[len(s.stack)-1]   // Get the last element
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
	fmt.Println("size of stack: ", len(s.stack))
}

// NewStack initializes and returns an empty stack
func NewStack() Stack {
	return Stack{
		stack: []int{},
	}
}
