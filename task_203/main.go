package main

func main() {
	l := &ListNode{Val: 1, Next: &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, nil}}}}}
	removeElements(l, 1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	if head == nil {
		return nil
	}

	current := head.Next
	prev := head
	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
			current.Next = nil
			current = prev
		}
		prev = current
		current = current.Next
	}
	return head
}
