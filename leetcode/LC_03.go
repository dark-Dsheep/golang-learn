package main

/*
*
3. 无重复字符的最长子串
*/
func lengthOfLongestSubstring(s string) (ans int) {
	mp := make(map[int]int)
	j := 0
	for i, x := range s {
		mp[int(x)]++
		for mp[int(x)] > 1 && j < len(s) {
			mp[int(s[j])]--
			j++
		}
		ans = max(ans, i-j+1)
		i++
	}
	return ans
}

func main() {

}
