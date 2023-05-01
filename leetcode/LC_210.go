package main

/*
*
210. 课程表 II
*/
func findOrder(n int, prerequisites [][]int) []int {
	rd := make([]int, n)
	g := make([][]int, n)
	for i := range g {
		g[i] = []int{}
	}
	for _, p := range prerequisites {
		x, y := p[0], p[1]
		g[x] = append(g[x], y)
		rd[y]++
	}
	q := []int{}
	for i := 0; i < n; i++ {
		if rd[i] == 0 {
			q = append(q, i)
		}
	}
	ans := []int{}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		ans = append(ans, p)
		for _, v := range g[p] {
			rd[v]--
			if rd[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(ans) != n {
		return []int{}
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

func main() {

}
