package main

import (
	"errors"
	"fmt"
	"math"
)

// 直接声明
func DoSomeThing() {
}

// 字面量
var DoSomething func()

// 类型
type DoAnyThing func()

func sum(a, b int) int {
	return a + b
}

// 同类型且相邻的参数 只需要声明一次类型
func max(a, b, c int) int {
	if a < b {
		a, b = b, a
	}
	if a < c {
		a, c = c, a
	}
	return a
}

// go的函数可以有多个返回值
func Div(a, b float64) (float64, error) {
	if a == 0 {
		return math.NaN(), errors.New("0不能作为被除数")
	}
	return a / b, nil
}

func main() {
	ans := sum(1, 2)
	fmt.Println(ans)
	mx := max(2, 4, 6)
	fmt.Println(mx)

	// 匿名函数
	func(a, b int) int {
		return a + b
	}(1, 2)

}
