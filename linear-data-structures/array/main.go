package main

import (
	"bufio"
	"fmt"
	"os"

	command "github.com/Data-Structures/array/command"
)

func main() {
	list := command.NewArrayList(bufio.NewReader(os.Stdin))

	for {
		fmt.Println(`
			Welcome to the Array List Program! This program allows you to perform the following operations:
				1. Insert an element
				2. Display the list
				3. Delete an element
				4. Search for an element
				5. Exit
			Please enter the number of the operation you would like to perform:
		`)

		var input int
		fmt.Scan(&input)
		list.HandleCases(int8(input))
	}
}
