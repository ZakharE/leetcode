package main

func main() {
	r := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}
	//2,3,3,4,5,null,4
	println(isSymmetric(r))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	return p.Val == q.Val && isSameTree(p.Right, q.Left) && isSameTree(p.Left, q.Right)
}
func isSymmetric(root *TreeNode) bool {
	return isSameTree(root.Left, root.Right)
}
