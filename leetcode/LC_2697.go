package main

// 2697. 字典序最小回文串
func makeSmallestPalindrome(S string) (ans string) {
	s := []byte(S)
	for i, n := 0, len(s); i < n/2; i++ {
		v, w := s[i], s[n-1-i]
		if v != w {
			if v > w {
				s[i] = w
			} else {
				s[n-1-i] = v
			}
		}
	}
	ans = string(s)
	return
}

func main() {

}
