package main

/*
*
1019. 链表中的下一个更大节点
*/

// ListNode /**
type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	a := make([]int, 10001)
	st := []int{}
	n, idx := 0, 0
	cur := head
	for cur != nil {
		n++
		a[idx] = cur.Val
		idx++
		cur = cur.Next
	}
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && st[0] <= a[i] {
			st = st[1:]
		}
		if len(st) > 0 {
			ans[i] = st[0]
		}
		st = append([]int{a[i]}, st...)
	}
	return ans
}

func main() {

}
