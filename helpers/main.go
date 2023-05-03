package main

import "fmt"

func main() {
	tree := generateTree([]any{1, 2, 2, 3, 3, nil, nil, 4, 4})
	fmt.Println(tree)
}

func generateTree(vals []any) *TreeNode {
	var root = &TreeNode{}
	toVisit := make([]*TreeNode, len(vals))
	//toVisit[0] = root
	v := 0
	for v < len(toVisit) {
		toVisit[v] = &TreeNode{}
		toVisit[v].Val = vals[v]
		//if v+1 < len(toVisit) {
		toVisit[v].Left = &TreeNode{}
		toVisit[v+1] = toVisit[v].Left
		//}
		//if v+2 < len(toVisit) {
		toVisit[v].Right = &TreeNode{}
		toVisit[v+2] = toVisit[v].Right

		v += 2
		//}
	}

	return root
}

type TreeNode struct {
	Val   any
	Left  *TreeNode
	Right *TreeNode
}
