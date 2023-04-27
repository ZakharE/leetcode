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
	println((r))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return depth(root, 0)
}

func depth(root *TreeNode, d int) int {
	if root == nil {
		return d
	}
	l := depth(root.Left, d+1)
	r := depth(root.Right, d+1)
	if l > r {
		return l
	}
	return r
}
