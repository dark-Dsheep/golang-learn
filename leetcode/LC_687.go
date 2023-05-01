package main

/*
*
687. 最长同值路径
*/
func longestUnivaluePath(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, v int) int {
		if root == nil {
			return -1
		}
		lmax := dfs(root.Left, root.Val) + 1
		rmax := dfs(root.Right, root.Val) + 1
		ans = max(ans, lmax+rmax)
		if root.Val != v {
			return -1
		}
		return max(lmax, rmax)
	}
	dfs(root, 0x3f3f3f)
	return ans
}

func main() {

}
