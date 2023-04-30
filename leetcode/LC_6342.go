package main

/*
*
6342. 找出叠涂元素
*/
func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	r, c := make(map[int]int), make(map[int]int)
	mp := make(map[int][]int)
	for i, row := range mat {
		for j := range row {
			mp[mat[i][j]] = []int{i, j}
		}
	}
	for i, v := range arr {
		p := mp[v]
		r[p[0]]++
		c[p[1]]++
		if r[p[0]] == n || c[p[1]] == m {
			return i
		}
	}
	return 0
}

func main() {

}
