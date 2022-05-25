package aencoding

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func encodingUsage(str string) {
	encoderGbkStr, err := simplifiedchinese.GBK.NewEncoder().String(str)
	if err != nil {
		fmt.Printf("err is : %v\n", err.Error())
	} else {
		// 这里是乱码，证明原来str是utf-8编码  注：go中的string均采用utf-8编码
		fmt.Printf("encoderGbkStr is: %v\n", encoderGbkStr)
	}

	gbkByte, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))
	utf8Byte, _ := simplifiedchinese.GBK.NewDecoder().Bytes(gbkByte)
	fmt.Printf("utf8Byte is: %s\n", string(utf8Byte))

	decoderGbkStr, _ := simplifiedchinese.GBK.NewDecoder().String(encoderGbkStr)
	fmt.Printf("decoderGbkStr is: %v\n", decoderGbkStr)

	gbkS := base64.StdEncoding.EncodeToString(gbkByte)
	decodeString, _ := base64.StdEncoding.DecodeString(gbkS)
	bytes, _ := simplifiedchinese.GBK.NewDecoder().Bytes(decodeString)

	fmt.Printf("gbkdebase64Str is : %s\n", string(bytes))

	utf8base64Str := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Printf("base64Str utf8 is: %s \n", utf8base64Str)
	i, _ := base64.StdEncoding.DecodeString(utf8base64Str)
	fmt.Printf("base64Decode is: %s\n", string(i))

}
