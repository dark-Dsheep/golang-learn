package main

/*
*
2423. 删除字符使频率相同
*/
func equalFrequency(word string) bool {
	count := make([]int, 200)
	for _, v := range word {
		count[v]++
	}
	var check func() bool
	check = func() bool {
		cnt := 0
		for i := 'a'; i <= 'z'; i++ {
			if count[i] != 0 {
				if cnt == 0 {
					cnt = count[i]
				} else if cnt != count[i] {
					return false
				}
			}
		}
		return true
	}
	for _, v := range word {
		count[v]--
		if check() {
			return true
		}
		count[v]++
	}
	return false
}

func main() {

}
