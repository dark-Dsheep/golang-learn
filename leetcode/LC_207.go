package main

/*
*
207. 课程表
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	rd := make([]int, numCourses) // 入度
	g := make([][]int, numCourses)
	for i := range g {
		g[i] = []int{}
	}
	for _, p := range prerequisites {
		x, y := p[0], p[1]
		g[x] = append(g[x], y)
		rd[y]++
	}
	q := []int{}
	for i := 0; i < numCourses; i++ {
		if rd[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		u := q[0] // pop
		q = q[1:]
		numCourses--
		for _, v := range g[u] {
			rd[v]--
			if rd[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return numCourses == 0
}

func main() {
	// ...
}
