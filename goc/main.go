package main

/*

#cgo CFLAGS: -Iinclude

#cgo LDFLAGS: -Llib -llibvideo

#include "video.h"

*/
import "C"

// 编译.c文件生成.so文件： gcc video.c -fPIC -shared -o libvideo.so
func main() {
	cmd := C.CString("ffmpeg -i ./xxx/*.png ./xxx/yyy.mp4")
	C.exeFFmpegCmd(&cmd)
}
