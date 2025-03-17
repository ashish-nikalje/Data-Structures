package command

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// InsertNode inserts a new node with user-provided data at the end of the list.
func (l *LinkedList) InsertNode() {
	var data int
	fmt.Print("enter the data to insert: ")
	fmt.Scan(&data)

	node := &Node{data: data}

	if l.head == nil {
		l.head = node
	} else {
		temp := l.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = node
	}

	fmt.Println("node inserted successfully.")
}

// DisplayList prints all elements of the linked list.
func (l *LinkedList) DisplayList() {
	if l.head == nil {
		fmt.Println("list is empty.")
		return
	}

	fmt.Print("linked list: ")
	temp := l.head
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
	fmt.Println("null")
}

// DeleteNode removes the first occurrence of a node with specified data.
func (l *LinkedList) DeleteNode() {
	if l.head == nil {
		fmt.Println("list is empty.")
		return
	}

	var data int
	fmt.Print("enter the data to delete: ")
	fmt.Scan(&data)

	if l.head.data == data {
		l.head = l.head.next
		fmt.Println("node deleted successfully.")
		return
	}

	temp := l.head
	for temp.next != nil {
		if temp.next.data == data {
			temp.next = temp.next.next
			fmt.Println("node deleted successfully.")
			return
		}
		temp = temp.next
	}

	fmt.Println("node not found.")
}

// SearchNode checks if a given data value exists in the list.
func (l *LinkedList) SearchNode() {
	if l.head == nil {
		fmt.Println("list is empty.")
		return
	}

	var data int
	fmt.Print("enter the data to search for: ")
	fmt.Scan(&data)

	temp := l.head
	for temp != nil {
		if temp.data == data {
			fmt.Println("node found.")
			return
		}
		temp = temp.next
	}

	fmt.Println("node not found.")
}
