package base64api

import (
	"encoding/base64"
	"fmt"
	"log"
)

func Encode(str string) string {
	src := []byte(str)
	dst := base64.StdEncoding.EncodeToString(src)
	fmt.Println("解码后的数据为：", dst)
	return dst
}

func Decode(str64 string) []byte {
	dst, err := base64.StdEncoding.DecodeString(str64)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("解码后的数据为：", string(dst))
	return dst
}
