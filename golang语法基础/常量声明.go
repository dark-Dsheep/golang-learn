package main

/**
常量的定义
*/

const mod = 1e9 + 7
const name string = "Jack"                // 字面量
const msg string = "hello world"          // 字面量
const num = 1                             // 字面量
const numExpression = (1+2+3)/2%100 + num // 常量表达式

// 下面的常量声明 编译无法通过
//const desc string

// 批量声明常量 使用()
const (
	Count = 1
	Desc  = "flow"
)

const (
	Size = 16
	Len  = 25
)

// iota 是一个内置的常量标识符
const iota = 0
const (
	Num = iota
	Num1
	Num2
	Num3
	Num4
)

// 常量值无法被修改
func main() {
	//Size++
}
