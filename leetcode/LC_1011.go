package main

/*
*1011. 在 D 天内送达包裹的能力
 */
func shipWithinDays(weights []int, days int) int {
	l, r := 0, 25000000
	for _, x := range weights {
		if x > l {
			l = x
		}
	}
	var check func(int) bool
	check = func(mid int) bool {
		d, sum := days-1, 0
		for _, x := range weights {
			if sum+x > mid {
				d--
				sum = 0
			}
			sum += x
		}
		return d >= 0
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
