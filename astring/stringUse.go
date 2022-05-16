package astring

import (
	"fmt"
	"strconv"
	"strings"
)

func StringUsage() {
	str := "The quick brown fox jumps over the lazy dog"
	fmt.Printf("输入字符串为：%s\n", str)
	// len()  index
	length := len(str)
	fmt.Printf("长度：%d\n", length)

	index := strings.Index(str, "quick")
	fmt.Printf("quick下标: %d\n", index)

	s1 := str[:index]
	s2 := str[index:]
	fmt.Printf("根据下标取值: %s, %s\n", s1, s2)

	// split用法
	split := strings.Split(str, " ")
	for _, va := range split {
		fmt.Printf("%s - ", va)
	}
	fmt.Println()

	// join用法
	join := strings.Join(split, "-")
	fmt.Printf("%s\n", join)

	// 带进制转换string
	parseInt2, _ := strconv.ParseInt("10", 2, 32)
	parseInt10, _ := strconv.ParseInt("10", 10, 32)
	parseInt16, _ := strconv.ParseInt("10", 16, 32)
	fmt.Printf("2进制解'10'为数字 %d\n", parseInt2)
	fmt.Printf("10进制解'10'为数字 %d\n", parseInt10)
	fmt.Printf("16进制解'10'为数字 %d\n", parseInt16)

	// 解析为布尔值
	parseBool, _ := strconv.ParseBool("true")
	fmt.Printf("'true'解为bool: %t", parseBool)

}
