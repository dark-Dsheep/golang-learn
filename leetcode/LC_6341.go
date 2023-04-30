package main

/*
*
6341. 保龄球游戏的获胜者
*/
func isWinner(a []int, b []int) int {
	var calc func([]int) int
	calc = func(arr []int) (s int) {
		s = arr[0]
		for i := 1; i < len(arr); i++ {
			if i >= 2 {
				if arr[i-1] == 10 || arr[i-2] == 10 {
					s += 2 * arr[i]
				} else {
					s += arr[i]
				}
			} else {
				if arr[i-1] == 10 {
					s += 2 * arr[i]
				} else {
					s += arr[i]
				}
			}
		}
		return s
	}
	s1, s2 := calc(a), calc(b)
	if s1 == s2 {
		return 0
	}
	if s1 > s2 {
		return 1
	} else {
		return 2
	}
}

func main() {

}
