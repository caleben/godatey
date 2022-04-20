package main

/*
  go调用C的代码示例
*/

/*
#include <stdio.h>

void printInt(int v) {
    printf("printInt: %d\n", v);
}
*/
import "C"

func main() {
	v := 42
	C.printInt(C.int(v))
}
