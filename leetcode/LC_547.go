package main

// 547. 省份数量
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	cnt := n
	fa := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fa[i] = i
	}
	var find func(int) int
	var merge func(int, int)
	find = func(x int) int {
		for x != fa[x] {
			fa[x] = fa[fa[x]]
			x = fa[x]
		}
		return x
	}
	merge = func(a int, b int) {
		a = find(a)
		b = find(b)
		if a == b {
			return
		}
		fa[b] = a
		cnt--
	}
	for i, con := range isConnected {
		for j, x := range con {
			if x == 1 {
				merge(i, j)
			}
		}
	}
	return cnt
}

func main() {

}
