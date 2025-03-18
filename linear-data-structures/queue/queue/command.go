package queue

import (
	"fmt"
	"os"
)

func (c *Queue) HandleCases(caseID int) {
	switch caseID {
	case 1:
		c.Enqueue()
		return
	case 2:
		c.DisplayQueue()
		return
	case 3:
		c.Dequeue()
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
