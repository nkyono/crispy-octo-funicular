package BinaryTree

import (
	"fmt"
	"math"
)

type BinaryTree struct {
	Val   int
	Right *BinaryTree
	Left  *BinaryTree
}

// Creates a binary tree using a list of integers
func CreateBinaryTree(values []int) *BinaryTree {
	var bt *BinaryTree

	for _, v := range values {
		bt.AddNode(v)
	}

	return bt
}

// Recursive binary tree comparison
func CompareTrees(a, b *BinaryTree) bool {
	// if both are nil than they are equivalent
	if a == nil && b == nil {
		return true
	}
	// if only one of them is nil than they are not equivalent
	if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val {
		// if both have the same value then return true if the rest of the trees are the same
		return CompareTrees(a.Left, b.Left) && CompareTrees(a.Right, b.Right)
	}
	return false
}

func (bt *BinaryTree) AddNode(val int) {

}

func (bt *BinaryTree) RemoveNode(val int) {

}

// checks to see whether a binary tree is balanced (returns <=1 if balanced)
// a binary tree is balanced if the left and right tree heights differ by no more than 1
// is a 'private' (go doesn't have private, but it is lowercase so it's pretty much private equivalent) and is only used for Binary Tree construction
func (bt *BinaryTree) checkBalance() bool {
	if bt == nil {
		return true
	}
	leftHeight := bt.Left.Height()
	rightHeight := bt.Right.Height()

	return (math.Abs(float64(leftHeight)-float64(rightHeight)) <= 1 && bt.Left.checkBalance() && bt.Right.checkBalance())
}

// returns the height of the binary tree
func (bt *BinaryTree) Height() int {
	if bt == nil {
		return 0
	}
	return int(math.Max(float64(bt.Left.Height()), float64(bt.Right.Height()))) + 1
}

// returns a string of the binary tree in preorder
func (bt *BinaryTree) ToString() string {
	treeString := ""
	if bt == nil {
		return treeString
	}
	stack := []*BinaryTree{}
	stack = append(stack, bt)
	for len(stack) != 0 {
		treeString = treeString + fmt.Sprint(stack[0].Val)
		if stack[0].Left != nil {
			stack = append(stack, stack[0].Left)
		}
		if stack[0].Right != nil {
			stack = append(stack, stack[0].Right)
		}
		stack = stack[1:]
	}
	return treeString
}
