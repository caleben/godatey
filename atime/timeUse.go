package atime

import (
	"fmt"
	"time"
)

const (
	layout         = "2006-01-02 15:04:05.000"
	layoutWithZone = "2006-01-02 15:04:05.000+0800"
)

func timeUsage(timeStr string) {
	input := timeStr
	fmt.Printf("输入时间字符串: %s\n", timeStr)

	// 带时区的parse
	localTime, err := time.ParseInLocation(layout, input, time.Local)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("字符串转化为Time: %v\n", localTime)
	format := localTime.Format(layoutWithZone)
	fmt.Printf("Time转化为字符串: %v\n", format)

	// 获取时间戳10位及13位
	fmt.Printf("unix时间戳(10): %d, unix时间戳(13): %d\n", localTime.Unix(), localTime.UnixNano()/1e6)

	// 时间戳转化为Time（默认时区为time.local）
	unixT := 1654086020
	unixTime := time.Unix(int64(unixT), 123e6)
	fmt.Printf("时间戳转为Time: %v\n", unixTime)

	// 本地时区转化为UTC
	utcTime := localTime.UTC()
	fmt.Printf("本地时间转化为UTC时间：%v\n", utcTime)

	// UTC时区转化为本地
	fmt.Printf("UTC时间转化为本地时间：%v\n", utcTime.Local())

}
