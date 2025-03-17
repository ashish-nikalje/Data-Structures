package command

import (
	"bufio"
	"fmt"
	"os"
)

type ArrayList struct {
	data []int
}

func NewArrayList(scanner *bufio.Reader) *ArrayList {
	return &ArrayList{data: []int{}}
}

func (a *ArrayList) HandleCases(caseID int8) {
	switch caseID {
	case 1:
		a.InsertElement()
	case 2:
		a.DisplayList()
	case 3:
		a.DeleteElement()
	case 4:
		a.SearchElement()
	case 5:
		fmt.Println("exiting program....")
		os.Exit(0)
	default:
		fmt.Println("yet to be implemented")
	}
}
