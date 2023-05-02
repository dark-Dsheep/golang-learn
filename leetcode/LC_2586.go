package main

/**
2586. 统计范围内的元音字符串数
*/

func vowelStrings(words []string, left int, right int) (ans int) {
	var check func(int) bool
	check = func(x int) bool {
		return x == 'a' || x == 'e' || x == 'i' || x == 'o' || x == 'u'
	}
	for i := left; i <= right; i++ {
		if check(int(words[i][0])) && check(int(words[i][len(words[i])-1])) {
			ans++
		}
	}
	return ans
}

func main() {

}
