package main

//337. 打家劫舍 III

func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	lRob, lNotRob := dfs(node.Left)
	rRob, rNotRob := dfs(node.Right)
	rob := lNotRob + rNotRob + node.Val
	notRob := max(lRob, lNotRob) + max(rRob, rNotRob)
	return rob, notRob
}

func rob3(root *TreeNode) int {
	return max(dfs(root))
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func main() {

}
