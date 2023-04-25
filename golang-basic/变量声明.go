package main

import "fmt"

/**
变量
*/

var intNum int
var str string
var char byte

// 一次性声明多个同类型变量
var x, y, z int

// 声明多个不同类型的变量 使用()进行包裹
var (
	username string
	age      int
	address  string
)

var (
	school string
	class  int
)

// 赋值操作
var ss = "jack"
var val = 100

func main() {
	// 使用 := 初始化变量
	cc := "cc"
	fmt.Println(cc)

	// 使用 := 批量初始化
	username, age := "jack", 1
	fmt.Println(username)
	fmt.Println(age)
	fmt.Println("=============================")
	// 交换变量值
	x1, x2, x3 := 1, 2, 3
	x1, x2, x3 = x3, x2, x1
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
}
