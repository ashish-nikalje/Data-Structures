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
		return
	case 4:
		c.SearchNode()
		return
	case 5:
		c.ClearList()
		return
	case 6:
		fmt.Println("exiting program....")
		os.Exit(0)
	default:
		fmt.Println("yet to be implemented")
	}
}
