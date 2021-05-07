package linkedList

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func (l *LinkNode) ToString() string {
	res := ""
	curr := l
	for curr != nil {
		res = res + fmt.Sprint(curr.Val)
		curr = curr.Next
	}
	return res
}

func (l *LinkNode) AddNode(newNode *LinkNode) {
	curr := l
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

// call head.ReverseLinkedList(), returns new head
func (l *LinkNode) ReverseLinkedList() *LinkNode {
	var prev *LinkNode
	curr := l
	for curr.Next != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	curr.Next = prev
	return curr
}

func (a *LinkNode) CompareLinkedList(b *LinkNode) bool {
	if a == nil && b == nil {
		return true
	}
	// after for loop, one or both should be nil
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	// check if a or b is nil
	return (a == nil && b == nil)
}
