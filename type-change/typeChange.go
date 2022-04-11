package main

import (
	"fmt"
	"reflect"
)

/**
 * 两种方式判断一个数据的真实类型
 */
func main() {
	var i interface{}
	i = "Hello"
	i = 1234
	// 利用ok-idiom模式将接口类型还原为原始类型
	if t, ok := i.(string); ok {
		fmt.Printf("%v World!\n", t)
	} else {
		fmt.Printf("need string, but %v\n", reflect.TypeOf(i).Name())
	}

	// 通过 type switch将接口类型还原为原始类型
	switch t := i.(type) {
	case string:
		fmt.Printf("%v World!\n", t)
	case int:
		fmt.Printf("%d 5678\n", t)
	default:
		fmt.Printf("error\n")
	}
}
