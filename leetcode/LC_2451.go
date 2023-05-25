package main

// 2451. 差值数组不同的字符串
func oddString(words []string) string {
	m := map[string][]string{}
	diff := make([]byte, len(words[0])-1)
	for _, s := range words {
		for i := 0; i < len(s)-1; i++ {
			diff[i] = s[i] - s[i+1]
		}
		t := string(diff)
		m[t] = append(m[t], s)
	}
	for _, g := range m {
		if len(g) == 1 {
			return g[0]
		}
	}
	return ""
}

func main() {

}
