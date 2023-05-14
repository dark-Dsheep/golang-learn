package main

/*
*6431. 相邻值的按位异或
 */
func doesValidArrayExist(derived []int) bool {
	cnt := 0
	for _, x := range derived {
		if x == 1 {
			cnt++
		}
	}
	return cnt%2 == 0
}

func main() {

}
