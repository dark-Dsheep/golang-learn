package main

import "math"

/*
*76. 最小覆盖子串
 */

func minWindow(s string, t string) string {
	//need := make([]int, 128)
	//for _, x := range t {
	//	need[x]++
	//}
	//j, count, st, length := 0, len(t), 0, math.MaxInt
	//for i, x := range s {
	//	if need[x] > 0 {
	//		count--
	//	}
	//	need[x]--
	//	if count == 0 {
	//		for j < i && need[s[j]] < 0 {
	//			need[s[j]]++
	//			j++
	//		}
	//		if length > i-j+1 {
	//			length = i - j + 1
	//			st = j
	//		}
	//		need[s[j]]++
	//		j++
	//		count++
	//	}
	//}
	//if length == math.MaxInt {
	//	return ""
	//}
	window, need := make([]int, 128), make([]int, 128)
	for _, x := range t {
		need[x]++
	}
	var getLen func() int
	getLen = func() (ans int) {
		for _, x := range need {
			if x != 0 {
				ans++
			}
		}
		return ans
	}
	j, count, st, minLen := 0, 0, 0, math.MaxInt
	for i, x := range s {
		window[x]++
		if window[x] == need[x] {
			count++
		}
		for count == getLen() {
			if i-j+1 < minLen {
				minLen = i - j + 1
				st = j
			}
			if window[s[j]] == need[s[j]] {
				count--
			}
			window[s[j]]--
			j++
		}
	}
	if minLen == math.MaxInt {
		return ""
	}
	return s[st : st+minLen]
}

func main() {
	minWindow("aa", "aa")
}
