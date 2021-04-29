package main

import "fmt"

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
	for len(stack) != 0 {
		fmt.Printf("%d ", stack[0].Val)
		if stack[0].Left != nil {
			stack = append(stack, stack[0].Left)
		}
		if stack[0].Right != nil {
			stack = append(stack, stack[0].Right)
		}
		stack = stack[1:]
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
	for _, v := range vals[1:] {
		insertNode(v, root)
	}
	return root
}

func main() {
	root := createTree([]int{4, 2, 7, 1, 3, 6, 9})
	printTree(root)
}
