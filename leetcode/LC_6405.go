package main

/*
*
6405. 找到两个数组的前缀公共数组
*/
func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	cnt := 0
	set1, set2 := make([]int, 100), make([]int, 100)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		x, y := A[i], B[i]
		if set2[x] != 0 {
			cnt++
		}
		if set1[y] != 0 {
			cnt++
		}
		if x == y {
			cnt++
		}
		set1[x]++
		set2[y]++
		ans[i] = cnt
	}
	return ans
}

func main() {

}
