package main

/*
*1760. 袋子里最少数目的球
 */
func minimumSize(nums []int, maxOperations int) int {
	var check func(int) bool
	check = func(mid int) bool {
		cnt := 0
		for _, x := range nums {
			cnt += (x - 1) / mid
		}
		return cnt <= maxOperations
	}
	l, r := 1, 1000000000
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
