package main

/**
128. 最长连续序列
*/

var fa = [100010]int{}

func find(x int) int {
	for x != fa[x] {
		fa[x] = fa[fa[x]]
		x = fa[x]
	}
	return x
}

func union(a, b int) {
	a, b = find(a), find(b)
	if a == b {
		return
	}
	fa[b] = a
}

func longestConsecutive(nums []int) (ans int) {
	mp := make(map[int]int)
	for i, v := range nums {
		fa[i] = i
		mp[v] = i
	}
	for i, v := range nums {
		if idx, ok := mp[v+1]; ok {
			union(i, idx)
		}
	}
	count := make(map[int]int)
	for _, v := range mp {
		root := find(v)
		count[root]++
	}
	for _, cnt := range count {
		ans = max(ans, cnt)
	}
	return ans
}
