package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Data-Structures/linear-data-structures/queue/queue"
)

func main() {
	list := queue.NewQueue()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(`
welcome to the queue program! this program allows you to perform the following operations on queue:
1. insert into the queue
2. display the queue
3. delete from the queue
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
