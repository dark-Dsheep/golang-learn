package main

import "sort"

/*
*2517. 礼盒的最大甜蜜度
 */
func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	l, r := 0, 1000000000
	var check func(int) bool
	check = func(mid int) bool {
		cnt, pre := 1, price[0]
		for _, x := range price[1:] {
			if x-pre >= mid {
				cnt++
				pre = x
			}
		}
		return cnt >= k
	}
	for l < r {
		mid := (l + r + 1) / 2
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func main() {

}
