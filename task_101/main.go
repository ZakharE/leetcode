package main

func main() {
	bst := sortedArrayToBST([]int{1, 2, 3})
	println(bst)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	middle := len(nums) / 2
	node := &TreeNode{Val: nums[middle]}

	node.Left = sortedArrayToBST(nums[0:middle])
	node.Right = sortedArrayToBST(nums[middle+1:])
	return node
}
