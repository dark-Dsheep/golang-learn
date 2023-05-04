package main

/*
*230. 二叉搜索树中第K小的元素
 */
func kthSmallest(root *TreeNode, k int) (ans int) {
	cnt := 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		cnt++
		if cnt == k {
			ans = root.Val
			return
		}
		dfs(root.Right)
	}
	dfs(root)
	return ans
}

func main() {

}
