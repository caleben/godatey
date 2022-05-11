package json

import (
	"fmt"
	"github.com/json-iterator/go"
)

func Parse(input string) {
	str := input
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	data := new(map[string]interface{})
	// 方式一： 这种解析是解析全部json string,其中这里的data是map，也可以是一个具体的struct结构休
	err := json.UnmarshalFromString(str, &data)
	if err == nil {
		fmt.Printf("%v\n", data)
	}

	// 方式二: 这种就只解析这个字段 mypackage.rep，适用于 “定向解析”
	info := json.Get([]byte(str), "source_node")
	// 这个接口来判断key是否存在
	if info.GetInterface() != nil {
		toString := info.Get("host").ToString()
		var m = make(map[string]string)
		m["host"] = toString
		fmt.Printf("host is: %v\n", toString)
	}

	host1 := json.Get([]byte(str), "source_node", "host").ToString()
	fmt.Println("host is: " + host1)

}
