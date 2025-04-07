package command

import (
	"container/list"
	"fmt"
)

type node struct {
	data      int
	leftNode  *node
	rightNode *node
}

// isNil checks if a given node is nil
func isNil(n *node) bool {
	return n == nil
}

// NewBinarySearchTree returns an empty binary search tree
func NewBinarySearchTree() *node {
	return nil
}

// newNode creates a new node with the given data
func newNode(data int) *node {
	return &node{data: data}
}

// insertNode inserts a new node into the BST using pointer to pointer
func insertNode(n **node, data int) {
	if isNil(*n) {
		*n = newNode(data)
		return
	}
	if data < (*n).data {
		insertNode(&((*n).leftNode), data)
	} else if data > (*n).data {
		insertNode(&((*n).rightNode), data)
	}
}

// searchNode searches for a node with the given data
func (n *node) searchNode(data int) *node {
	if isNil(n) || n.data == data {
		return n
	}
	if data < n.data {
		return n.leftNode.searchNode(data)
	}
	return n.rightNode.searchNode(data)
}

// displayTreeTopDown prints the tree structure from top to bottom
func (n *node) displayTreeTopDown(level int) {
	if isNil(n) {
		return
	}
	n.rightNode.displayTreeTopDown(level + 1)
	for i := 0; i < level; i++ {
		fmt.Print("     ")
	}
	fmt.Println(n.data)
	n.leftNode.displayTreeTopDown(level + 1)
}

// inorderTraversal prints the BST in inorder (LNR)
func (n *node) inorderTraversal() {
	if isNil(n) {
		return
	}
	n.leftNode.inorderTraversal()
	fmt.Printf("%d ", n.data)
	n.rightNode.inorderTraversal()
}

// preorderTraversal prints the BST in preorder (NLR)
func (n *node) preorderTraversal() {
	if isNil(n) {
		return
	}
	fmt.Printf("%d ", n.data)
	n.leftNode.preorderTraversal()
	n.rightNode.preorderTraversal()
}

// postorderTraversal prints the BST in postorder (LRN)
func (n *node) postorderTraversal() {
	if isNil(n) {
		return
	}
	n.leftNode.postorderTraversal()
	n.rightNode.postorderTraversal()
	fmt.Printf("%d ", n.data)
}

// levelOrderTraversal prints the BST level by level (BFS)
func (n *node) levelOrderTraversal() {
	if isNil(n) {
		return
	}

	queue := list.New()
	queue.PushBack(n)

	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)
		current := elem.Value.(*node)
		fmt.Printf("%d ", current.data)

		if !isNil(current.leftNode) {
			queue.PushBack(current.leftNode)
		}
		if !isNil(current.rightNode) {
			queue.PushBack(current.rightNode)
		}
	}
}
