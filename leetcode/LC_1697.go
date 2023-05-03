package main

import "sort"

/**
1697. 检查边长度限制的路径是否存在
*/

//var fa = [100010]int{}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})
	index := make([]int, len(queries))
	for i := range queries {
		index[i] = i
	}
	sort.Slice(index, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	ans := make([]bool, len(queries))
	i := 0
	for _, query := range index {
		for i < len(edgeList) && edgeList[i][2] < queries[query][2] {
			union(edgeList[i][0], edgeList[i][1])
			i++
		}
		ans[query] = connected(queries[query][0], queries[query][1])
	}
	return ans
}

//func find(x int) int {
//	for x != fa[x] {
//		fa[x] = fa[fa[x]]
//		x = fa[x]
//	}
//	return x
//}
//
//func union(a, b int) {
//	a, b = find(a), find(b)
//	if a == b {
//		return
//	}
//	fa[b] = a
//}

func connected(a, b int) bool {
	return find(a) == find(b)
}

func main() {

}
