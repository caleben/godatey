package base64api

import (
	"encoding/base64"
	"fmt"
)

func Encode(str string) string {
	src := []byte(str)
	dst := base64.RawStdEncoding.EncodeToString(src)
	fmt.Println("加密后的数据为：" + dst)
	return dst
}

func Decode(str64 string) ([]byte, error) {
	decodeString, err := base64.RawStdEncoding.DecodeString(str64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("解密后的数据：" + string(decodeString))
	return decodeString, err
}
