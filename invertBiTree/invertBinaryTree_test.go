package test

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	i := 0
	for len(stack) != 0 {
		fmt.Printf("%d ", stack[i].Val)
		if stack[i].Left != nil {
			stack = append(stack, stack[i].Left)
		}
		if stack[i].Right != nil {
			stack = append(stack, stack[i].Right)
		}
		stack = stack[1:]
		i++
	}
}

func insertNode(val int, root *TreeNode) {
	curr := root
	var prev *TreeNode
	newNode := &TreeNode{val, nil, nil}
	for curr != nil {
		if curr.Val < val {
			if curr.Right != nil {
				prev = curr
				curr = curr.Right
			} else {
				curr.Right = newNode
				return
			}
		} else if curr.Val > val {
			if curr.Left != nil {
				prev = curr
				curr = curr.Left
			} else {
				curr.Left = newNode
				return
			}
		}
	}
	if val > prev.Val {
		prev.Right = newNode
	} else {
		prev.Left = newNode
	}
}

func createTree(vals []int) *TreeNode {
	if len(vals) <= 0 {
		return nil
	}
	root := &TreeNode{vals[0], nil, nil}
	for _, v := range vals {
		insertNode(v, root)
	}
	return root
}

func swapTrees(root *TreeNode) {
	if root == nil {
		return
	}
	// there may be an issue with a nil leaf
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
}

// Leetcode: faster than 100%, uses less memory than 6.58% (I would assume because of recursion)
// invertTreeRecursive, invertTreeRecursiveII, and invertTreeIterative all use the same memory soooo..... idk
func invertTreeRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	swapTrees(root)
	invertTreeRecursive(root.Left)
	invertTreeRecursive(root.Right)

	return root
}

func invertTreeRecursiveII(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	right := invertTreeRecursiveII(root.Left)
	left := invertTreeRecursiveII(root.Right)
	root.Right = right
	root.Left = left
	return root
}

func inverTreeIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		temp := curr.Left
		curr.Left = curr.Right
		curr.Right = temp
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return root
}

func compSliceInt(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestInvertBinaryTree(t *testing.T) {
	root := createTree([]int{4, 2, 7, 1, 3, 6, 9})
	printTree(root)
}
