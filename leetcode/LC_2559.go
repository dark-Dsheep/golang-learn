package main

import "strings"

/*
*
2559. 统计范围内的元音字符串数
*/
func vowelStrings2(words []string, queries [][]int) []int {
	n := len(words)
	sum := make([]int, n+1)
	for i, w := range words {
		if strings.Contains("aeiou", w[:1]) && strings.Contains("aeiou", w[len(w)-1:]) {
			sum[i+1] = sum[i] + 1
		} else {
			sum[i+1] = sum[i]
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = sum[q[1]+1] - sum[q[0]]
	}
	return ans
}

func main() {

}
