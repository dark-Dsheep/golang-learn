package main

import "fmt"

/**
42. 接雨水
*/

func trap(height []int) (ans int) {
	max, idx := 0, 0
	for i, v := range height {
		if v > max {
			idx = i
			max = v
		}
	}
	i, j := 0, len(height)-1
	lmax, rmax := 0, 0
	// ==================left==================
	for i < idx {
		if height[i] > lmax {
			lmax = height[i]
		} else {
			ans += lmax - height[i]
		}
		i++
	}
	// ==================right=================
	for j > idx {
		if height[j] > rmax {
			rmax = height[j]
		} else {
			ans += rmax - height[j]
		}
		j--
	}
	return ans
}

func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	ans := trap(height)
	fmt.Println(ans)
}
