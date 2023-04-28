package main

import (
	"fmt"
	"strings"
)

/**
strings库
*/

func main() {

	// 比较字符串
	fmt.Println(strings.Compare("abc", "abe"))
	fmt.Println(strings.Compare("ab", "ab"))

	// 包含字符串
	fmt.Println(strings.Contains("abcdefg", "abc"))
	fmt.Println(strings.Contains("abc", "ec"))

	// 子串出现次数
	fmt.Println(strings.Count("this is a str", "s"))
	fmt.Println(strings.Count("3.1415926", "1"))

	// 删除指定子串
	s := "aabcbv"
	before, after, found := strings.Cut(s, "bc")
	fmt.Println(before, after, found)

	// 忽略大小写
	fmt.Println(strings.EqualFold("Hello", "hello"))

	// 分割字符串 返回 []string
	fmt.Printf("%q\n", strings.Fields(" a b c d e f g"))                       // 根据空格分割
	fmt.Printf("%q\n", strings.FieldsFunc("a,b,c,d,e,f,g", func(r rune) bool { // 根据返回值分割
		return r == ','
	}))

	// 寻找前后缀
	str := "abbc cbba"
	fmt.Println(strings.HasPrefix(str, "abb"))
	fmt.Println(strings.HasSuffix(str, "bba"))

	// 子串的位置 Index
	fmt.Println(strings.Index("abcdefg", "bc"))
	fmt.Println(strings.IndexAny("abcdefg", "cb")) // 任意子串下标
	fmt.Println(strings.IndexRune("abcdefg", 'g'))
	// 最后一次出现的字符串下标
	fmt.Println(strings.LastIndex("abcfgfg", "fg"))

	// 替换字符串 n->替换次数
	fmt.Println(strings.Replace("Hello this is golang", "golang", "c++", 1))
	fmt.Println(strings.Replace("Hello this is golang", "o", "c", 2))

	// 分隔字符串
	fmt.Printf("%q\n", strings.Split("this is go language", " "))
	fmt.Printf("%q\n", strings.SplitN("this is go language", " ", 2)) // n->分隔次数

	// 大小写转换
	fmt.Println(strings.ToLower("My Name is"))
	fmt.Println(strings.ToUpper("you are wawawa"))

	// 修剪字符串 trim
	fmt.Println(strings.Trim(" this is a test str ", " "))
}
