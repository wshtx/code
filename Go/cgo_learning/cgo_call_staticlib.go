package main

//#cgo CFLAGS: -I./include
//#cgo LDFLAGS: -L${SRCDIR}/lib -lmymath
//#include "mymath.h"
import "C"
import "fmt"

func TestCallStaticLib(a, b int) (int, error) {
	res, err := C.add(C.int(a), C.int(b))
	if err != nil {
		return 0, err
	}
	fmt.Println("Test call static lib：调用外部库的文件属于main包，则在main函数所在文件或者该文件选择一个指定编译链接参数即可")
	return int(res), nil
}
