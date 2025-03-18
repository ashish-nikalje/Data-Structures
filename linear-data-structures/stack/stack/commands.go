package stack

import (
	"fmt"
	"os"
)

func (c *Stack) HandleCases(caseID int) {
	switch caseID {
	case 1:
		c.Push()
		return
	case 2:
		c.Pop()
		return
	case 3:
		c.DisplayStack()
		return
	case 4:
		c.Peek()
		return
	case 5:
		c.Size()
		return
	case 6:
		fmt.Println("exiting program....")
		os.Exit(0)
	default:
		fmt.Println("yet to be implemented")
	}
}
