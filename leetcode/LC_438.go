package main

/*
*
438. 找到字符串中所有字母异位词
*/
func findAnagrams(s string, p string) (ans []int) {
	scnt := make([]int, 26)
	pcnt := make([]int, 26)
	for _, x := range p {
		pcnt[x-'a']++
	}
	var check func() bool
	check = func() bool {
		for i := 0; i < 26; i++ {
			if scnt[i] != pcnt[i] {
				return false
			}
		}
		return true
	}
	j := 0
	for i, x := range s {
		scnt[x-'a']++
		if i-j+1 == len(p) {
			if check() {
				ans = append(ans, j)
			}
			scnt[s[j]-'a']--
			j++
		}
		i++
	}
	return ans
}

func main() {

}
