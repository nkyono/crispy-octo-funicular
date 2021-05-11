package BinaryTree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	var testBinaryTreeCreation = []struct {
		in   []int
		want *BinaryTree
	}{
		{[]int{1, 2, 3}, &BinaryTree{2, &BinaryTree{1, nil, nil}, &BinaryTree{3, nil, nil}}},
		{[]int{1}, &BinaryTree{1, nil, nil}},
		{[]int{2, 1, 4, 5, 3}, &BinaryTree{2, &BinaryTree{1, nil, nil}, &BinaryTree{4, &BinaryTree{3, nil, nil}, &BinaryTree{5, nil, nil}}}},
	}

	for _, tt := range testBinaryTreeCreation {
		testname := fmt.Sprintf("%v", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := CreateBinaryTree(tt.in)
			if !CompareTrees(ans, tt.want) {
				t.Errorf("got %s, want %s", ans.ToString(), tt.want.ToString())
			}
		})
	}
}
