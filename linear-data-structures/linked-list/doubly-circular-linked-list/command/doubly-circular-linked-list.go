package command

import "fmt"

type LinkedList struct {
	head *Node // pointer to the first node
}

// Node represents each element of the list
type Node struct {
	Data int   // data stored in the node
	next *Node // pointer to the next node
	prev *Node // pointer to the previous node
}

// InsertNode adds a new node at the end of the circular doubly linked list
func (l *LinkedList) InsertNode() {
	var data int
	print("enter the data to insert: ")
	fmt.Scan(&data)

	node := &Node{Data: data}

	if l.isEmpty() {
		l.head = node    // set the head to the new node
		node.next = node // link to itself
		node.prev = node // link to itself
		return
	}

	tail := l.head.prev // get the tail node

	tail.next = node   // link tail to new node
	node.prev = tail   // link new node back to tail
	node.next = l.head // link new node to head
	l.head.prev = node // link head back to new node

	fmt.Println("node inserted successfully.")
}

// DisplayList prints the elements of the circular doubly linked list
func (l *LinkedList) DisplayList() {
	if l.isEmpty() {
		fmt.Println("list is empty.")
		return
	}

	temp := l.head
	for {
		fmt.Printf("[%d]", temp.Data)
		temp = temp.next
		if temp == l.head {
			break // looped back to start
		}
		fmt.Print(" <-> ")
	}

	fmt.Printf(" -> [%d] (head)\n", l.head.Data) // print the head node
}

// DeleteNode removes the first occurrence of a node with the given data
func (l *LinkedList) DeleteNode() {
	if l.isEmpty() {
		fmt.Println("list is empty.")
		return
	}

	var data int
	print("enter the data to delete: ")
	fmt.Scan(&data)

	if l.head.Data == data { // case 1: deleting the head node
		if l.head.next == l.head { // only one node in the list
			l.head = nil // set head to nil
		} else { // more than one node in the list
			tail := l.head.prev  // get the tail node
			l.head = l.head.next // move head to next node
			tail.next = l.head   // link tail to new head
			l.head.prev = tail   // link new head back to tail
		}

		fmt.Println("node deleted successfully.")
		return
	}

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

	fmt.Println("node not found.")
}

// SearchNode searches for a node with the specified data
func (l *LinkedList) SearchNode() {
	if l.isEmpty() {
		fmt.Println("list is empty.")
		return
	}

	var data int
	print("enter the data to search for: ")
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

// ClearList clears the entire list
func (l *LinkedList) ClearList() {
	if l.isEmpty() {
		fmt.Println("list is already empty.")
		return
	}

	l.head = nil
	fmt.Println("list cleared successfully.")
}

// ReverseList reverses the circular doubly linked list
func (l *LinkedList) ReverseList() {
	if l.isEmpty() {
		fmt.Println("list is empty.")
		return
	}

	fmt.Println("list before reversal: ")
	l.DisplayList()

	var prev *Node
	current := l.head

	for {
		next := current.next
		current.next = prev
		current.prev = next

		prev = current
		current = next

		if current == l.head {
			break // looped back to start
		}
	}

	l.head.next = prev // link head to new tail
	prev.prev = l.head // link new tail back to head

	fmt.Println("list reversed successfully.")
	l.DisplayList()
}

func (l *LinkedList) isEmpty() bool {
	return l.head == nil
}
