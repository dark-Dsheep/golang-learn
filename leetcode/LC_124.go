package main

import "math"

/*
*
124. 二叉树中的最大路径和
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lVal := dfs(root.Left)
		rVal := dfs(root.Right)
		ans = max(ans, lVal+rVal+root.Val)
		return max(max(lVal, rVal)+root.Val, 0)
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {

}
