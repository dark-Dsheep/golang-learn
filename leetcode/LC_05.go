package main

/**
5. 最长回文子串
*/

func longestPalindrome(s string) string {
	n, maxLen := len(s), 1
	ans := s
	for i := 1; i < n; i++ {
		x := s[i]
		l, r := i, i
		for l >= 0 && s[l] == x {
			l--
		}
		for r < n && s[r] == x {
			r++
		}
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l+1 > maxLen {
			ans = s[l+1 : r]
			maxLen = r - l + 1
		}
	}
	return ans
}

func main() {

}
