package cgo_call_lib

//#cgo CFLAGS: -I../include
//#cgo LDFLAGS: -L${SRCDIR}/lib -lmymath
//#include "mymath.h"
import "C"
import "fmt"

func TestCallStaticLib(a, b int) (int, error) {
	res, err := C.add(C.int(a), C.int(b))
	if err != nil {
		return 0, err
	}
	fmt.Println("Test call static lib：非main包中的文件使用外部库时，需要在main包和非main包同时指定编译连接参数（CFLAGS/LDFLAGS）")
	return int(res), nil
}
