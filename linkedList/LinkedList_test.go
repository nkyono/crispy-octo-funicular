package linkedList

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	var testAdd = []struct {
		in   *LinkNode
		want string
	}{
		{&LinkNode{1, nil}, "01"},
		{&LinkNode{2, nil}, "012"},
		{&LinkNode{3, nil}, "0123"},
		{&LinkNode{4, nil}, "01234"},
	}

	l := &LinkNode{0, nil}
	for _, tt := range testAdd {
		testname := fmt.Sprint(tt.want)
		t.Run(testname, func(t *testing.T) {
			l.AddNode(tt.in)
			res := l.ToString()
			if res != tt.want {
				t.Errorf("got %s, want %s", res, tt.want)
			}
		})
	}

	var testReverse = []struct {
		in   *LinkNode
		want *LinkNode
	}{
		{&LinkNode{1, &LinkNode{2, &LinkNode{3, nil}}}, &LinkNode{3, &LinkNode{2, &LinkNode{1, nil}}}},
		{&LinkNode{1, &LinkNode{2, &LinkNode{2, nil}}}, &LinkNode{2, &LinkNode{2, &LinkNode{1, nil}}}},
		{&LinkNode{1, &LinkNode{2, nil}}, &LinkNode{2, &LinkNode{1, nil}}},
		{&LinkNode{1, nil}, &LinkNode{1, nil}},
	}

	for _, tt := range testReverse {
		testname := fmt.Sprint(tt.in.ToString())
		t.Run(testname, func(t *testing.T) {
			ans := tt.in.ReverseLinkedList().ToString()
			wantString := tt.want.ToString()
			if ans != wantString {
				t.Errorf("got %s, want %s", ans, wantString)
			}
		})
	}

	var testCompare = []struct {
		inA  *LinkNode
		inB  *LinkNode
		want bool
	}{
		{&LinkNode{1, &LinkNode{2, &LinkNode{3, nil}}}, &LinkNode{3, &LinkNode{2, &LinkNode{1, nil}}}, false},
		{&LinkNode{2, &LinkNode{2, &LinkNode{2, nil}}}, &LinkNode{2, &LinkNode{2, &LinkNode{2, nil}}}, true},
		{&LinkNode{1, nil}, &LinkNode{2, nil}, false},
		{&LinkNode{1, nil}, &LinkNode{1, nil}, true},
		{&LinkNode{1, nil}, &LinkNode{1, &LinkNode{2, &LinkNode{2, nil}}}, false},
		{&LinkNode{1, &LinkNode{2, &LinkNode{3, nil}}}, &LinkNode{1, nil}, false},
		{&LinkNode{1, &LinkNode{2, &LinkNode{3, nil}}}, nil, false},
	}

	for _, tt := range testCompare {
		testname := fmt.Sprint(tt.inA.ToString())
		t.Run(testname, func(t *testing.T) {
			ans := tt.inA.CompareLinkedList(tt.inB)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

	var testMerge = []struct {
		inA  *LinkNode
		inB  *LinkNode
		want *LinkNode
	}{
		{&LinkNode{1, &LinkNode{2, &LinkNode{3, nil}}}, &LinkNode{4, &LinkNode{5, nil}}, &LinkNode{1, &LinkNode{2, &LinkNode{3, &LinkNode{4, &LinkNode{5, nil}}}}}},
		{&LinkNode{1, nil}, &LinkNode{2, nil}, &LinkNode{1, &LinkNode{2, nil}}},
		{&LinkNode{1, nil}, nil, &LinkNode{1, nil}},
		{nil, &LinkNode{2, nil}, &LinkNode{2, nil}},
	}

	for _, tt := range testMerge {
		testname := fmt.Sprintf("Merge %s", tt.inA.ToString())
		t.Run(testname, func(t *testing.T) {
			ans := MergeList(tt.inA, tt.inB)
			if tt.want.CompareLinkedList(ans) {
				t.Errorf("got %v, want %v", ans.ToString(), tt.want.ToString())
			}
		})
	}

}
