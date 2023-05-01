package main

/*
*875. 爱吃香蕉的珂珂
 */
func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 1000000000
	var check func(int) bool
	check = func(mid int) bool {
		cost := 0
		for _, x := range piles {
			cost += (x-1)/mid + 1
		}
		return cost <= h
	}
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {

}
