package main

/**
2554. 从一个范围内选择最多整数 I
*/

func maxCount(banned []int, n int, maxSum int) int {
	set := make(map[int]bool)
	for _, x := range banned {
		set[x] = true
	}
	ans, sum := 0, 0
	for i := 1; i <= n; i++ {
		if sum+i > maxSum {
			break
		}
		if !set[i] {
			sum += i
			ans++
		}
	}
	return ans
}

func main() {

}
