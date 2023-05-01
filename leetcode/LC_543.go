package main

/*
*
543. 二叉树的直径
*/
func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return -1
		}
		lmax := dfs(root.Left) + 1
		rmax := dfs(root.Right) + 1
		ans = max(ans, lmax+rmax)
		return max(lmax, rmax)
	}
	dfs(root)
	return ans
}

func main() {

}
