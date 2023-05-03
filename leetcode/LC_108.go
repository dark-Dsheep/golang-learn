package main

/*
*
108. 将有序数组转换为二叉搜索树
*/
func sortedArrayToBST(nums []int) *TreeNode {

	var dfs func(int, int) *TreeNode
	dfs = func(l int, r int) *TreeNode {
		if l > r {
			return nil
		}
		mid := (l + r) / 2
		root := &TreeNode{Val: nums[mid]}
		root.Left = dfs(l, mid-1)
		root.Right = dfs(mid+1, r)
		return root
	}
	return dfs(0, len(nums)-1)
}

func main() {

}
