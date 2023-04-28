package main

import (
	"fmt"
	"math"
)

// 最大值
func TestMax() {
	fmt.Println(math.Max(1.0, 2.0))
}

// 最小值
func TestMin() {
	fmt.Println(math.Min(1.1, 2.2))
}

// 绝对值
func TestAbs() {
	fmt.Println(math.Abs(-1))
}

// 余数
func TestMod() {
	fmt.Println(math.Mod(1, 10))
	fmt.Println(math.Mod(12, 10))
}

// 取整
func TestTrunc() {
	fmt.Println(math.Trunc(1.26))
	fmt.Println(math.Trunc(2.3333))
}

// 下取整
func TestFloor() {
	fmt.Println(math.Floor(2.5))
}

// 上取整
func TestCeil() {
	fmt.Println(math.Ceil(2.5))
}

// 四舍五入
func TestRound() {
	fmt.Println(math.Round(1.2389))
	fmt.Println(math.Round(-5.2389))
}

// 求对数
func TestLog() {
	fmt.Println(math.Log(100) / math.Log(10))
	fmt.Println(math.Log(1) / math.Log(2))
}

// E的指数
func TestEx() {
	fmt.Println(math.Exp(2))
}

// 幂
func TestPow() {
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Pow(3, 3))
}

// 平方根
func TestSqrt() {
	fmt.Println(math.Sqrt(4))
}

// 立方根
func TestCube() {
	fmt.Println(math.Cbrt(3))
}

// 开N方
func TestN() {
	fmt.Println(math.Round(math.Pow(8, 1.0/3)))
	fmt.Println(math.Round(math.Pow(100, 1.0/2)))
}

func main() {
	TestMax()
	TestMin()
	TestAbs()
	TestMod()
	TestTrunc()
	//TestFloor()
}
