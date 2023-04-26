package main

/*
*
410. 分割数组的最大值
*/
func splitArray(nums []int, k int) int {
	l, r := 0, 0
	for _, x := range nums {
		r += x
		if l < x {
			l = x
		}
	}
	var check func([]int, int, int) bool
	check = func(nums []int, mid int, k int) bool {
		s := 0
		for _, x := range nums {
			if s+x > mid {
				k--
				s = 0
			}
			s += x
		}
		return k > 0
	}

	for l < r {
		mid := (r-l)/2 + l
		if check(nums, mid, k) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {

}
