package command

import (
	"bufio"
	"fmt"
	"os"
)

func NewLinkedList(scanner *bufio.Reader) *LinkedList {
	return &LinkedList{}
}

func (c *LinkedList) HandleCases(caseID int8) {
	switch caseID {
	case 1:
		c.InsertNode()
	case 2:
		c.DisplayList()
	case 3:
		c.DeleteNode()
	case 4:
		c.SearchNode()
	case 5:
		fmt.Println("exiting program....")
		os.Exit(0)
	default:
		fmt.Println("yet to be implemented")
	}
}
