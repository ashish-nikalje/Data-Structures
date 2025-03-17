package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	command "github.com/Data-Structures/linear-data-structures/linked-list/single-linked-list/command"
)

func main() {
	list := command.NewLinkedList(bufio.NewReader(os.Stdin))
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(`
welcome to the single linked list program! this program allows you to perform the following operations:
1. insert a node
2. display the list
3. delete a node
4. search for a node
5. exit
please enter the number of the operation you would like to perform:`)

		scanner.Scan()
		input, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("invalid input, please enter a number between 1 and 5")
			continue
		}

		list.HandleCases(int8(input))
	}
}
