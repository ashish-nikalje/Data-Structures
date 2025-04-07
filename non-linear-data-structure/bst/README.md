# Binary Search Tree (BST)

## Overview
This project contains a practical and interactive implementation of a **Binary Search Tree (BST)** in Go. It allows you to insert and search elements, view the tree structure, and explore various types of traversals to understand how BSTs work.

Whether you're brushing up on data structures or building a tool from scratch, this module provides a solid foundation.

## Features
- Insert elements into the tree
- Search for a specific value
- Display the tree visually (top-down)
- Traverse the tree:
  - In-order (LNR)
  - Pre-order (NLR)
  - Post-order (LRN)
  - Level-order (BFS)

## File Structure
```
bst/
├── command/
│   ├── cmd.go         # Handles user interaction and menu logic
│   └── node.go        # Contains BST impl and helper func.
├── go.mod             # Go module definition
├── main.go            # Application entry point
└── README.md          # Documentation file (this one)
```

## How to Run
Make sure you're inside the `bst` directory and run:
```bash
go run main.go
```

You’ll see a friendly menu like this:

```
Welcome to the Binary Search Tree Program!
Choose an operation:
1. Insert an element
2. Search for an element
3. Display Tree (Top-Down)
4. Inorder Traversal
5. Preorder Traversal
6. Postorder Traversal
7. Level Order Traversal
8. Exit
```

## Example Traversal Outputs

Assume you've inserted: `50, 30, 70, 20, 40, 60, 80`

- **In-order (LNR)**:  
  `20 30 40 50 60 70 80`

- **Pre-order (NLR)**:  
  `50 30 20 40 70 60 80`

- **Post-order (LRN)**:  
  `20 40 30 60 80 70 50`

- **Level-order (BFS)**:  
  `50 30 70 20 40 60 80`



## Tips
- You can only insert **unique**, **non-negative** integers.