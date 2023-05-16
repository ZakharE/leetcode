package main

import "strconv"

func main() {
	root := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
		Left: &TreeNode{
			Val: 3,
		},
	}
	binaryTreePaths(&root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return make([]string, 0)
	}
	res := make([]string, 0, 100)
	str := strconv.Itoa(root.Val)

	buildStr(root.Right, str, &res)
	buildStr(root.Left, str, &res)
	return res
}

func buildStr(root *TreeNode, str string, res *[]string) {
	str = str + "->" + strconv.Itoa(root.Val)
	if root == nil || root.Left == nil && root.Right == nil {
		*res = append(*res, str)
		return
	}
	buildStr(root.Left, str, res)
	buildStr(root.Right, str, res)
}
