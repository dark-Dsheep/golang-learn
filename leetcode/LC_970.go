package main

import "math"

/*
*
970. 强整数
*/
func powerfulIntegers(x int, y int, bound int) (ans []int) {
	set := map[int]struct{}{}
	for i := 0; i < 21; i++ {
		v1 := pow(x, i)
		if v1 > bound {
			break
		}
		for j := 0; j < 21; j++ {
			sum := v1 + pow(y, j)
			if sum > bound {
				break
			}
			if _, ok := set[sum]; !ok {
				set[sum] = struct{}{}
				ans = append(ans, sum)
			}
		}
	}
	return ans
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {

}
