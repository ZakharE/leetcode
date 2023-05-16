package main

func main() {
	println(isPalindrome(&ListNode{1, &ListNode{Val: 1, Next: &ListNode{2, &ListNode{Val: 1}}}}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	firstHalf := head
	secondHalf := reverseList(findMiddle(head))
	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}
	return true
}

func findMiddle(head *ListNode) *ListNode {
	slow := head
	var fast *ListNode
	if slow.Next == nil {
		return slow
	}
	if slow.Next.Next == nil {
		return slow.Next
	}
	fast = slow.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow.Next
}

func reverseList(head *ListNode) *ListNode {
	var prev, current, next *ListNode
	current = head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
