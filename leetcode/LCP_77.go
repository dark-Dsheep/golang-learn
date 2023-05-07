package main

import "sort"

/**
LCP 77. 符文储备
*/

func runeReserve(runes []int) int {
	sort.Ints(runes)
	ans, pre := 0, 1
	for i := 1; i < len(runes); i++ {
		if runes[i]-runes[i-1] > 1 {
			ans = max(ans, pre)
			pre = 1
		} else {
			pre++
		}
	}
	if pre != 0 {
		ans = max(ans, pre)
	}
	return ans
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func main() {

}
