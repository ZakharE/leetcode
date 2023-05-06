package main

func main() {
	//[3,2,0,-4]
	n := &ListNode{
		Val:  0,
		Next: &ListNode{Val: 4},
	}
	l := &ListNode{
		Val:  3,
		Next: n,
	}
	n.Next = l
	//l := &ListNode{Val: 3,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: n,
	//}
	hasCycle(l)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	for slow, fast := head, head; fast != nil && fast.Next != nil; {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
