package main

/*
*6418. 有相同颜色的相邻元素数目
 */
func colorTheArray(n int, queries [][]int) []int {
	ans := make([]int, len(queries))
	a := make([]int, n)
	i, sum := 0, 0
	for _, q := range queries {
		idx, c := q[0], q[1]
		if a[idx] != 0 {
			if idx-1 >= 0 && idx+1 < n {
				if a[idx-1] == a[idx] && a[idx+1] == a[idx] {
					sum -= 2
				} else if a[idx-1] == a[idx] || a[idx+1] == a[idx] {
					sum -= 1
				}
			} else if idx-1 >= 0 && a[idx-1] == a[idx] {
				sum -= 1
			} else if idx+1 < n && a[idx+1] == a[idx] {
				sum -= 1
			}
		}
		a[idx] = c
		if idx-1 >= 0 && idx+1 < n {
			if a[idx-1] == a[idx] && a[idx+1] == a[idx] {
				sum += 2
			} else if a[idx-1] == a[idx] || a[idx+1] == a[idx] {
				sum += 1
			}
		} else if idx-1 >= 0 && a[idx-1] == a[idx] {
			sum += 1
		} else if idx+1 < n && a[idx+1] == a[idx] {
			sum += 1
		}
		ans[i] = sum
		i++
	}
	return ans
}

func main() {

}
