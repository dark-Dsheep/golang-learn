package main

import "fmt"

func main() {

	//mp := map[int]string{
	//	0: "a",
	//	1: "a",
	//	2: "a",
	//	3: "a",
	//	4: "a",
	//}

	//var mp map[string]int
	//mp["a"] = 1
	mp := make(map[string]int, 8)
	mp["a"] = 1
	//mp2 := make(map[string][]int, 10)
	fmt.Println(mp)

	// 判断key是否存在
	if val, f := mp["b"]; f {
		fmt.Println(val)
	} else {
		fmt.Println("key不存在")
	}
	// 删除
	delete(mp, "a")

}
