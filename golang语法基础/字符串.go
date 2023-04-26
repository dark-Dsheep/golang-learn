package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "this is a string"
	fmt.Println(str[0])
	// 切割字符串
	fmt.Println(str[0:4])
	// 修改字符串的字符
	//str[0] = 'a'    编译无法通过
	str = "that is a string"
	fmt.Println(str)

	// 转换
	ss := "this is a string"
	// 显示类型转换为字节切片
	bytes := []byte(ss)
	// 显示类型转换为字符串
	fmt.Println(string(bytes))

	s1 := "abc"
	s2 := "aaa"
	fmt.Println(len(s1), len(s2))

	// 获取其中一个字符
	fmt.Println(string(s1[0]))

	// 高性能的字符串拼接 strings.Builder
	sb := strings.Builder{}
	sb.WriteString("abc")
	sb.WriteString("def")
	fmt.Println(sb.String())

}
