package main

//var dirs = []struct{ x, y int }{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

/*
*994. 腐烂的橘子
 */
func orangesRotting(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	q := [][]int{}
	cnt := 0 // 好橘子数目
	for i, row := range grid {
		for j, v := range row {
			if v == 2 {
				q = append(q, []int{i, j})
			} else if v == 1 {
				cnt++
			}
		}
	}
	if cnt == 0 {
		return 0
	}
	for len(q) > 0 {
		sz := len(q)
		for sz > 0 {
			p := q[0]
			q = q[1:] // pop
			x, y := p[0], p[1]
			for _, d := range dirs {
				x, y := x+d.x, y+d.y
				if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 2 || grid[x][y] == 0 {
					continue
				}
				grid[x][y] = 2
				q = append(q, []int{x, y})
				cnt--
			}
			sz--
		}
		ans++
	}
	if cnt != 0 {
		return -1
	}
	return ans - 1
}

func main() {

}
