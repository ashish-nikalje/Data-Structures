package command

import (
	"fmt"
)

// LinkedList represents the circular doubly linked list
type LinkedList struct {
	head *Node
}

// Node represents each element of the list
type Node struct {
	Data int   // data stored in the node
	next *Node // pointer to the next node
	prev *Node // pointer to the previous node
}

// insertNode adds a new node at the end of the circular doubly linked list
func (l *LinkedList) InsertNode() {
	var data int
	fmt.Print("enter the data to insert: ")
	fmt.Scan(&data)

	// create new node with the given data
	newNode := &Node{Data: data}

	// if the list is empty, point the node to itself (circular)
	if l.isEmpty() { // check if the list is empty
		newNode.next = newNode // point to itself
		newNode.prev = newNode // point to itself
		l.head = newNode       // set head to new node
		return
	}

	// otherwise, insert at the end
	tail := l.head.prev // last node before head

	tail.next = newNode   // tail -> newNode
	newNode.prev = tail   // newNode <- tail
	newNode.next = l.head // newNode -> head
	l.head.prev = newNode // head <- newNode

	fmt.Println("node inserted successfully.")
}

// displayList prints the elements of the circular doubly linked list
func (l *LinkedList) DisplayList() {
	if l.isEmpty() {
		fmt.Println("list is empty.")
		return
	}

	fmt.Print("list: ")
	temp := l.head

	// loop through the circular list
	for {
		fmt.Printf("[%d]", temp.Data)
		temp = temp.next
		if temp == l.head {
			break
		}
		fmt.Print(" -> ")
	}
	fmt.Printf(" -> [%d] (head)\n", l.head.Data) // shows it's circular
}

// deleteNode removes the first occurrence of a node with the given data
func (l *LinkedList) DeleteNode() {
	if l.isEmpty() {
		fmt.Println("list is empty.")
		return
	}

	var data int
	fmt.Print("enter the data to delete: ")
	fmt.Scan(&data)

	// case 1: head node is to be deleted
	if l.head.Data == data {
		if l.head.next == l.head {
			// only one node in the list
			l.head = nil
			fmt.Println("node deleted successfully.")
			return
		}
		// more than one node
		tail := l.head.prev
		tail.next = l.head.next // tail -> second node
		l.head.next.prev = tail // second node <- tail
		l.head = l.head.next    // move head forward
		fmt.Println("node deleted successfully.")
		return
	}

	// case 2: deleting a non-head node
	temp := l.head.next
	for temp != l.head {
		if temp.Data == data {
			temp.prev.next = temp.next // bypass current node forward
			temp.next.prev = temp.prev // bypass current node backward
			fmt.Println("node deleted successfully.")
			return
		}
		temp = temp.next
	}

	// node not found after full circle
	fmt.Println("node not found.")
}

// searchNode searches for a node with the specified data
func (l *LinkedList) SearchNode() {
	if l.isEmpty() {
		fmt.Println("list is empty.")
		return
	}

	var data int
	fmt.Print("enter the data to search for: ")
	fmt.Scan(&data)

	temp := l.head
	for {
		if temp.Data == data {
			fmt.Println("node found.")
			return
		}
		temp = temp.next
		if temp == l.head {
			break // looped back to start
		}
	}
	fmt.Println("node not found.")
}

// clearList clears the entire list
func (l *LinkedList) ClearList() {
	if l.isEmpty() {
		fmt.Println("list is already empty.")
		return
	}

	l.head = nil
	fmt.Println("list cleared successfully.")
}

// isEmpty checks if the list is empty
func (l *LinkedList) isEmpty() bool {
	return l.head == nil
}
