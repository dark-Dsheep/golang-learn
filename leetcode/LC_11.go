package main

/*
*
11. 盛最多水的容器
*/
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	ans := 0
	for i < j {
		if height[i] > height[j] {
			ans = max(ans, height[j]*(j-i))
			j--
		} else {
			ans = max(ans, height[i]*(j-i))
			i++
		}
	}
	return ans
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func main() {

}
