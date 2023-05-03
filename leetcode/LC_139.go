package main

/*
*139. 单词拆分
 */
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	f := make([]bool, n+1)
	set := make(map[string]struct{})
	for _, w := range wordDict {
		set[w] = struct{}{}
	}
	f[0] = true
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := set[s[j:i]]; ok {
				f[i] = f[i] || f[j]
			}
		}
	}
	return f[n]
}

func main() {

}
