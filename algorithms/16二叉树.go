package algorithms

import "fmt"

type Tree struct {
	data  int
	left  *Tree
	right *Tree
}

func NewTreeNode(data int) *Tree {
	return &Tree{
		data,
		nil,
		nil,
	}
}

func preOrder(tree *Tree) {
	if tree != nil {
		fmt.Printf("%v-->", tree.data)
		preOrder(tree.left)
		preOrder(tree.right)
	}
}
func midOrder(tree *Tree) {
	if tree != nil {
		preOrder(tree.left)
		fmt.Printf("%v-->", tree.data)
		preOrder(tree.right)
	}
}
func afterOrder(tree *Tree) {
	if tree != nil {
		preOrder(tree.left)
		preOrder(tree.right)
		fmt.Printf("%v-->", tree.data)

	}
}

func treeTest() {
	root := NewTreeNode(0)
	root.left = NewTreeNode(1)
	root.right = NewTreeNode(2)
	root.left.left = NewTreeNode(3)
	root.left.right = NewTreeNode(4)
	root.right.left = NewTreeNode(5)
	root.right.right = NewTreeNode(6)

	preOrder(root)
	fmt.Println()
	midOrder(root)
	fmt.Println()
	afterOrder(root)
}
