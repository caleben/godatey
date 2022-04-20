package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	baseCount := 1000
	times := flag.Int("n", 1, "日志文件")
	inputFile := flag.String("c", "TraceLog.log", "日志文件")
	outputFile := flag.String("o", "outTrace.log", "日志文件")
	flag.Parse()
	fmt.Printf("inputFile is: %v\n", *inputFile)
	buf, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		panic(err.Error())
	}

	_, err = os.Stat(*outputFile)
	if err != nil {
		fmt.Println("文件不存在,创建文件...")
		_, _ = os.Create(*outputFile)
	}

	file, err := os.OpenFile(*outputFile, os.O_WRONLY|os.O_APPEND, 0x666)
	if err != nil {
		panic(err.Error())
	}
	writer := bufio.NewWriter(file)
	count := baseCount * (*times)
	for i := 0; i < count; i++ {
		_, err := writer.Write(buf)
		if err != nil {
			panic(err.Error())
		}
	}
	err1 := writer.Flush()
	if err1 != nil {
		fmt.Println("flush error: ", err.Error())
	}
	fmt.Printf("%d 文件已写完！\n", count)

}
