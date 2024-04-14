package main

import (
	"fmt"
	"prob404/tree"
)

func main() {
	fmt.Println("404 문제 풀이 -> https://leetcode.com/problems/sum-of-left-leaves/description/")
	// ex1 -> [3,9,20,null,null,15,7] => 9+15 = 24
	// ex2 -> [1] => 0
	// 1. -1000 <= val <= 1000
	// 2. The number of nodes in the tree is in the range [1, 1000].
	ex1 := &tree.TreeNode{3, &tree.TreeNode{9, nil, nil}, &tree.TreeNode{20, &tree.TreeNode{15, nil, nil}, &tree.TreeNode{7, nil, nil}}}
	fmt.Println(sumOfLeftLeaves(ex1)) // 24
}

func sumOfLeftLeaves(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	answer := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		answer += root.Left.Val
	}
	answer += sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	return answer
}
