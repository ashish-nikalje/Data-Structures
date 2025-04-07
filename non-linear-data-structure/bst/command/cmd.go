package command

import (
	"fmt"
	"os"
)

func HandleCases(n **node, caseID int) {
	switch caseID {
	case 1:
		var data int
		fmt.Print("insert: ")
		_, err := fmt.Scanf("%d", &data)
		if err != nil || data < 0 {
			fmt.Println("invalid input")
			return
		}
		insertNode(n, data)
		return

	case 2:
		var data int
		fmt.Print("search: ")
		_, err := fmt.Scanf("%d", &data)
		if err != nil || data < 0 {
			fmt.Println("invalid input")
			return
		}

		if (*n).searchNode(data) != nil {
			fmt.Println("found")
		} else {
			fmt.Println("not found")
		}
		return

	case 3:
		fmt.Println("tree:")
		(*n).displayTreeTopDown(0)
		return

	case 4:
		fmt.Print("inorder: ")
		(*n).inorderTraversal()
		fmt.Println()
		return

	case 5:
		fmt.Print("preorder: ")
		(*n).preorderTraversal()
		fmt.Println()
		return

	case 6:
		fmt.Print("postorder: ")
		(*n).postorderTraversal()
		fmt.Println()
		return

	case 7:
		fmt.Print("level order: ")
		(*n).levelOrderTraversal()
		fmt.Println()
		return

	case 8:
		fmt.Println("bye!")
		os.Exit(0)
		return

	default:
		fmt.Println("invalid choice")
		return
	}
}
