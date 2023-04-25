package main

/**
2643. 一最多的行
*/

func rowAndMaximumOnes(mat [][]int) []int {
	ans := make([]int, 2)
	mx := 0
	for i, row := range mat {
		cnt := 0
		for _, x := range row {
			if x == 1 {
				cnt++
			}
		}
		if cnt > mx {
			ans[0] = i
			ans[1] = cnt
			mx = cnt
		}
	}
	return ans
}

func main() {

}
