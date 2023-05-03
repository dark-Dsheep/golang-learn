package main

/*
*
1870. 准时到达的列车最小时速
*/
func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	var check func(int) bool
	check = func(mid int) bool {
		cost := float64(0)
		for _, x := range dist[0 : n-1] {
			cost += float64((x + mid - 1) / mid)
		}
		cost += float64(dist[n-1]) / float64(mid)
		return hour >= cost
	}
	if hour <= float64(n-1) {
		return -1
	}
	l, r := 1, 10000000
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
