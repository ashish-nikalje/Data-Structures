package command

import "fmt"

type LinkedList struct {
	head *Node
}

type Node struct {
	data int
	next *Node
	prev *Node
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
		node.prev = temp
	}

	fmt.Println("node inserted successfully.")
}

// DisplayList prints all elements of the linked list.
func (l *LinkedList) DisplayList() {
	if l.head == nil {
		fmt.Println("list is empty.")
		return
	}

	temp := l.head    // start from the head
	for temp != nil { // loop until the end of the list
		previousNode, nextNode := "null", "null" // default values for previous and next nodes
		if temp.prev != nil {                    // if the previous node exists, update the value
			previousNode = fmt.Sprintf("%d", temp.prev.data) // convert the data to a string
		}

		if temp.next != nil { // if the next node exists, update the value
			nextNode = fmt.Sprintf("%d", temp.next.data) // convert the data to a string
		}

		fmt.Printf("[%s <- %d -> %s] <-> ", previousNode, temp.data, nextNode)
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
		l.head.prev = nil
	}

	temp := l.head
	for temp.next != nil && temp.data != data { // loop until the end of the list or the data is found
		temp = temp.next // move to the next node
	}

	if temp.next == nil { // if the end of the list is reached
		fmt.Println("node not found.")
		return
	}

	if temp.next != nil {
		temp.next.prev = temp.prev
	}

	if temp.prev != nil {
		temp.prev.next = temp.next
	}

	fmt.Println("node deleted successfully.")
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
	for temp != nil && temp.data != data { // loop until the end of the list or the data is found
		temp = temp.next // move to the next node
	}

	if temp == nil { // if the end of the list is reached
		fmt.Println("node not found.")
		return
	}

	previousNode, nextNode := "null", "null" // default values for previous and next nodes
	if temp.prev != nil {                    // if the previous node exists, update the value
		previousNode = fmt.Sprintf("%d", temp.prev.data) // convert the data to a string
	}

	if temp.next != nil { // if the next node exists, update the value
		nextNode = fmt.Sprintf("%d", temp.next.data) // convert the data to a string
	}

	fmt.Printf("node found:  [%s <- %d -> %s]", previousNode, temp.data, nextNode)
}
