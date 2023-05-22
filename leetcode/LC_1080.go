package main

/*
*
1080. 根到叶路径上的不足节点
*/
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	var dfs func(*TreeNode, int) *TreeNode
	dfs = func(root *TreeNode, sum int) *TreeNode {
		if root == nil {
			return nil
		}
		sum += root.Val
		// 根到叶子节点的路径和
		if root.Left == nil && root.Right == nil {
			if sum < limit {
				return nil
			}
			return root
		}
		root.Left = dfs(root.Left, sum)
		root.Right = dfs(root.Right, sum)
		// 如果左右子树都是不足的, 那么根也返回null
		if root.Left == nil && root.Right == nil {
			return nil
		}
		return root
	}
	return dfs(root, 0)
}

func main() {

}
