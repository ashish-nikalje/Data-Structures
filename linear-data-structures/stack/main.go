package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Data-Structures/linear-data-structures/stack/stack"
)

func main() {
	list := stack.NewStack()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(`
welcome to the stack program! this program allows you to perform the following operations on stack:
1. push
2. pop
3. delete
4. peek
5. size
6. exit
please enter the number of the operation you would like to perform:`)

		scanner.Scan()
		input, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("invalid input, please enter a number between 1 and 5")
			continue
		}

		list.HandleCases(input)
	}
}
