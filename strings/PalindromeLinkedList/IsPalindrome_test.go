package PalindromeLinkedList

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// I think in theory this should have worked, but apparently it breaks with extremely large inputs
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	if head.Next == nil {
		return true
	}
	length := 0
	curr := head
	var prev *ListNode
	for curr != nil {
		length++
		prev = curr
		curr = curr.Next
	}
	if head.Val != prev.Val {
		return false
	}
	curr = head
	for curr.Next != prev {
		curr = curr.Next
	}
	curr.Next = nil
	head = head.Next

	return isPalindrome(head)
}

func (l *ListNode) ToString() string {
	res := ""
	curr := l
	for curr != nil {
		res = res + fmt.Sprint(curr.Val)
		curr = curr.Next
	}
	return res
}

func ReverseLinkedList(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	var prev *ListNode
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

func isPalindromeII(head *ListNode) bool {
	if head == nil {
		return true
	}
	if head.Next == nil {
		return true
	}
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}
	curr = head
	mid := (length + 1) / 2
	length = 0
	for curr != nil && length < mid {
		length++
		curr = curr.Next
	}
	headTwo := ReverseLinkedList(curr)

	for headTwo != nil && head != nil {
		if headTwo.Val != head.Val {
			return false
		}
		headTwo = headTwo.Next
		head = head.Next
	}

	return true
}

func TestIsPalindrome(t *testing.T) {
	var test = []struct {
		in   *ListNode
		want bool
	}{
		{&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}, true},
		{&ListNode{1, &ListNode{2, nil}}, false},
		{&ListNode{1, &ListNode{1, nil}}, true},
		{&ListNode{2, &ListNode{2, &ListNode{2, nil}}}, true},
		{&ListNode{1, &ListNode{2, &ListNode{1, nil}}}, true},
		{&ListNode{1, &ListNode{2, &ListNode{2, nil}}}, false},
		{&ListNode{2, &ListNode{2, &ListNode{1, nil}}}, false},
		{&ListNode{1, nil}, true},
	}

	for _, tt := range test {
		testname := fmt.Sprint(tt.in.ToString())
		t.Run(testname, func(t *testing.T) {
			ans := isPalindrome(tt.in)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

	var testTwo = []struct {
		in   *ListNode
		want bool
	}{
		{&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}, true},
		{&ListNode{1, &ListNode{2, nil}}, false},
		{&ListNode{1, &ListNode{1, nil}}, true},
		{&ListNode{2, &ListNode{2, &ListNode{2, nil}}}, true},
		{&ListNode{1, &ListNode{2, &ListNode{1, nil}}}, true},
		{&ListNode{1, &ListNode{2, &ListNode{2, nil}}}, false},
		{&ListNode{2, &ListNode{2, &ListNode{1, nil}}}, false},
		{&ListNode{1, nil}, true},
	}

	for _, tt := range testTwo {
		testname := fmt.Sprintf("Palindrome II: %s", tt.in.ToString())
		t.Run(testname, func(t *testing.T) {
			ans := isPalindromeII(tt.in)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
