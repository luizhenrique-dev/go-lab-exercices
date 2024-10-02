package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}

	fmt.Print("Original tree: ")
	fmt.Println(printTree(root))

	invertTree(root)

	fmt.Print("Inverted tree: ")
	fmt.Println(printTree(root))
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = root.Left, root.Right
	invertTree(root.Right)
	invertTree(root.Left)

	return root
}

func printTree(root *TreeNode) []int {
	var result []int
	if root != nil {
		result = append(result, root.Val)
		result = append(result, printTree(root.Left)...)
		result = append(result, printTree(root.Right)...)
	}

	return result
}
