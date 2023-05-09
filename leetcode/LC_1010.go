package main

/**
1010. 总持续时间可被 60 整除的歌曲
*/

func numPairsDivisibleBy60(time []int) (ans int) {
	mp := make([]int, 60)
	for _, x := range time {
		ans += mp[(60-x%60)%60]
		mp[x%60]++
	}
	return ans
}

func main() {

}
