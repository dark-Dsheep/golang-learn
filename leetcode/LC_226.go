package main

/*
*
226. 翻转二叉树
*/
func invertTree(root *TreeNode) *TreeNode {

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		temp := root.Left
		root.Left, root.Right = root.Right, temp
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return root
}

func main() {

}
