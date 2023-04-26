package main

/**
978. 最长湍流子数组
*/

func maxTurbulenceSize(arr []int) int {
	f := make([]int, 2)
	f[0], f[1] = 1, 1
	ans := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			f[0], f[1] = f[1]+1, 1
		} else if arr[i] < arr[i-1] {
			f[1], f[0] = f[0]+1, 1
		} else {
			f[0], f[1] = 1, 1
		}
		ans = max(ans, max(f[0], f[1]))
	}
	return ans
}

func main() {

}
