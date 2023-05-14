package main

/*
*
6430. 找出转圈游戏输家
*/
func circularGameLosers(n int, k int) (ans []int) {
	count := make(map[int]int)
	cur, base := 0, 1
	count[cur]++
	for count[cur] < 2 {
		cur += base * k
		if cur >= n {
			cur %= n
		}
		count[cur]++
		base++
	}
	for i := 1; i <= n; i++ {
		if count[i-1] == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {

}
