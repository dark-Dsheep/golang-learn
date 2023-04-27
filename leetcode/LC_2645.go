package main

/**
2645. 构造有效字符串的最少插入数
*/

func addMinimum(word string) int {
	t := 1
	for i := 1; i < len(word); i++ {
		if word[i-1] >= word[i] {
			t++
		}
	}
	return t*3 - len(word)
}

func main() {

}
