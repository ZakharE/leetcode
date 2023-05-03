package main

func main() {
	//[1,2,2,3,null,null,3,4,null,null,4]
	//1
	//2 2
	//3 n n 3
	//4 n   n 4
	t := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Right: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 15},
				Left:  &TreeNode{Val: 7},
			},
		},
	}
	balanced := isBalanced(t)
	println(balanced)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	res, _ := dfs(root)
	return res
}

func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	isLeftBalanced, leftHeight := dfs(root.Left)
	isRightBalanced, rightHeight := dfs(root.Right)
	diff := abs(leftHeight - rightHeight)
	if isLeftBalanced && isRightBalanced && diff <= 1 {
		return true, 1 + max(leftHeight, rightHeight)
	}
	return false, -1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
