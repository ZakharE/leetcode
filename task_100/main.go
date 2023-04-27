package main

func main() {
	r := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	l := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	tree := isSameTree(&r, &l)
	println(tree)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
