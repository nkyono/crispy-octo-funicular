package BinaryTree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	var testBinaryTreeComp = []struct {
		inTreeA *BinaryTree
		inTreeB *BinaryTree
		want    bool
	}{
		{&BinaryTree{1, nil, nil}, &BinaryTree{1, nil, &BinaryTree{2, nil, nil}}, false},
		{&BinaryTree{1, nil, nil}, &BinaryTree{1, nil, nil}, true},
		{nil, nil, true},
		{&BinaryTree{1, nil, &BinaryTree{2, nil, nil}}, &BinaryTree{1, nil, &BinaryTree{2, nil, nil}}, true},
		{&BinaryTree{1, nil, &BinaryTree{2, nil, nil}}, &BinaryTree{1, nil, &BinaryTree{2, nil, &BinaryTree{3, nil, nil}}}, false},
		{&BinaryTree{3, &BinaryTree{2, nil, nil}, &BinaryTree{4, nil, &BinaryTree{5, nil, nil}}}, &BinaryTree{3, &BinaryTree{2, &BinaryTree{1, nil, nil}, nil}, &BinaryTree{4, nil, &BinaryTree{5, nil, nil}}}, false},
		{&BinaryTree{3, &BinaryTree{2, &BinaryTree{1, nil, nil}, nil}, &BinaryTree{4, nil, &BinaryTree{5, nil, nil}}}, &BinaryTree{3, &BinaryTree{2, &BinaryTree{1, nil, nil}, nil}, &BinaryTree{4, nil, &BinaryTree{5, nil, nil}}}, true},
	}

	// something is wrong with my with CompareTrees(a,b) function
	// both print out the same string with ToString() but they aren't equivalent
	for _, tt := range testBinaryTreeComp {
		testname := fmt.Sprintf("%s %s", tt.inTreeA.ToString(), tt.inTreeB.ToString())
		t.Run(testname, func(t *testing.T) {
			ans := CompareTrees(tt.inTreeA, tt.inTreeB)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

	var testBinaryTreeAddNode = []struct {
		inTree *BinaryTree
		inVal  int
		want   *BinaryTree
	}{
		{&BinaryTree{1, nil, nil}, 2, &BinaryTree{1, nil, &BinaryTree{2, nil, nil}}},
		{&BinaryTree{1, nil, &BinaryTree{2, nil, nil}}, 3, &BinaryTree{1, nil, &BinaryTree{2, nil, &BinaryTree{3, nil, nil}}}},
		{&BinaryTree{3, &BinaryTree{2, nil, nil}, &BinaryTree{4, nil, &BinaryTree{5, nil, nil}}}, 1, &BinaryTree{3, &BinaryTree{2, &BinaryTree{1, nil, nil}, nil}, &BinaryTree{4, nil, &BinaryTree{5, nil, nil}}}},
	}

	for _, tt := range testBinaryTreeAddNode {
		testname := fmt.Sprintf("%s %d", tt.inTree.ToString(), tt.inVal)
		t.Run(testname, func(t *testing.T) {
			tt.inTree.AddNode(tt.inVal)
			if !CompareTrees(tt.inTree, tt.want) {
				t.Errorf("got %s, want %s", tt.inTree.ToString(), tt.want.ToString())
			}
		})
	}

	// test the full creation of a binary tree
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
