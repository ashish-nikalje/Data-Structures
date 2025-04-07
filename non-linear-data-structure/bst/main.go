package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	bst "github.com/Data-Structures/non-linear-data-structure/bst/command"
)

func main() {
	list := bst.NewBinarySearchTree()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(`
welcome to the stack program! this program allows you to perform the following operations on a binary search tree:
1. insert
2. search
2. display
4. inorder
5. preorder
6. postorder
7. level order
8. exit
please enter the number of the operation you would like to perform:`)

		scanner.Scan()
		input, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("invalid input, please enter a number between 1 and 5")
			continue
		}

		bst.HandleCases(&list, input)
	}
}
