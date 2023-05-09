package main

import (
	"fmt"
	"strconv"
)

// strconv库
//包 strconv 实现与基本数据类型的字符串表示形式之间的转换

func main() {
	// 字符串转整型
	ints, err := strconv.Atoi("456789")
	fmt.Println(ints, err)

	// 整型转字符串
	str := strconv.Itoa(114)
	fmt.Println(str)

	// 字符串转布尔值
	parseBool, err := strconv.ParseBool("1")
	fmt.Println(parseBool, err)

	b, err := strconv.ParseBool("true")
	fmt.Println(b, err)

	// 字符串转浮点数
	float, err := strconv.ParseFloat("1.145514", 64)
	fmt.Println(float, err)
}
