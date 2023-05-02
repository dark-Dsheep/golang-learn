package main

/*
*
102. 二叉树的层序遍历
*/
func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		list := make([]int, n)
		for i := 0; i < n; i++ {
			p := q[0]
			q = q[1:] // pop
			list[i] = p.Val
			if p.Left != nil {
				q = append(q, p.Left)
			}
			if p.Right != nil {
				q = append(q, p.Right)
			}
		}
		ans = append(ans, list)
	}
	return ans
}

func main() {

}
