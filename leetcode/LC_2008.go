package main

/*
*
2008. 出租车的最大盈利
*/
func maxTaxiEarnings(n int, rides [][]int) (ans int64) {
	f := make([]int, n+1)
	mp := make([][][2]int, n+1)
	for _, r := range rides {
		st, ed, tp := r[0], r[1], r[2]
		mp[ed] = append(mp[ed], [2]int{st, tp})
	}
	for end := 1; end <= n; end++ {
		f[end] = f[end-1]
		for _, r := range mp[end] {
			st, tp := r[0], r[1]
			f[end] = max(f[end], f[st]+end-st+tp)
		}
	}
	return int64(f[n])
}

func main() {

}
