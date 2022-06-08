package json

import (
	"fmt"
	"github.com/json-iterator/go"
	"strconv"
)

func Parse(input string) {
	str := input
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	data := new(map[string]interface{})
	arrStr := make([]interface{}, 10)
	var err error
	// 方式一： 这种解析是解析全部json string,其中这里的data是map，也可以是一个具体的struct结构休
	if input[0] == '{' {
		err = json.UnmarshalFromString(str, &data)
		if err == nil {
			fmt.Printf("%v\n", data)
			objToString, _ := jsoniter.MarshalToString(data)
			fmt.Printf("%s\n", objToString)
		} else {
			fmt.Printf("%v\n", err.Error())
		}
	} else if input[0] == '[' {
		err = json.UnmarshalFromString(str, &arrStr)
		if err == nil {
			fmt.Printf("%v\n", arrStr)
			arrToString, _ := jsoniter.MarshalToString(arrStr)
			fmt.Printf("%s\n", arrToString)
			for i, v := range arrStr {
				fmt.Printf("%T\n%v\n", v, v)
				toInt := jsoniter.Get([]byte(input), i, "_id").ToInt()
				fmt.Printf("-- %v\n", toInt)
			}
		} else {
			fmt.Printf("%v\n", err.Error())
		}
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

func changeJson() {
	input := `{"a":["b","c"], "status":404.00}`
	fmt.Printf("%s\n", input)
	var m1 map[string]interface{}
	err := jsoniter.UnmarshalFromString(input, &m1)
	if err != nil {
		return
	}
	arr := m1["a"]

	if t, ok := arr.([]interface{}); ok {
		t[0] = "bbb"
	} else {
		fmt.Println("not no no")
	}
	i := m1["status"]
	if a, ok := i.(float64); ok {
		fmt.Printf("%T-%v\n", i, strconv.FormatFloat(a, 'f', -1, 64))
	}
	out, _ := jsoniter.MarshalToString(m1)
	fmt.Printf("%s\n", out)
}
