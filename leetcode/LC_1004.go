package main

//1004. 最大连续1的个数 III

func longestOnes(nums []int, k int) (ans int) {
	f := make(map[int]int)
	maxCnt := 0
	for i, j := 0, 0; i < len(nums); i++ {
		f[nums[i]]++
		if nums[i] == 1 {
			maxCnt = max(maxCnt, f[nums[i]])
		}
		if i-j+1 > maxCnt+k {
			f[nums[j]]--
			j++
		}
		ans = max(ans, i-j+1)
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
