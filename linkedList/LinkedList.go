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

func MergeList(a, b *LinkNode) *LinkNode {
	if a == nil {
		return b
	} else if b == nil {
		return a
	}

	var ans *LinkNode
	if a.Val < b.Val {
		ans = a
	} else {
		ans = b
	}

	for a != nil && b != nil {
		if a.Val < b.Val {
			temp := a
			a = a.Next
			ans.Next = temp
		} else {
			temp := b
			b = b.Next
			ans.Next = temp
		}
		ans = ans.Next
	}

	if a != nil {
		ans.Next = a
	}
	if b != nil {
		ans.Next = b
	}

	return ans
}
