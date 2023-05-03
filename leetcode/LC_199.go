package main

/*
*
199. 二叉树的右视图
*/
func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			if i == n-1 {
				ans = append(ans, q[0].Val)
			}
			if q[0].Left != nil {
				q = append(q, q[0].Left)
			}
			if q[0].Right != nil {
				q = append(q, q[0].Right)
			}
			q = q[1:]
		}
	}
	return ans
}

func main() {

}
