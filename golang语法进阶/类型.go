package main

import "fmt"

// 类型声明
type MyInt int64
type MyFloat64 float64
type MyMap map[string]int

// 二维map
type TwoDMap = map[string]map[string]int

func PrintMyMap(mymap TwoDMap) {
	fmt.Println(mymap)
}

func main() {
	// 类型转换
	var x int8 = 1
	var y int32 = 512
	fmt.Println(int32(x), int8(y))
}
