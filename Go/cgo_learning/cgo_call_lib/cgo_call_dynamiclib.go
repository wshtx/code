package cgo_call_lib

//#cgo CFLAGS: -I../include
//#cgo LDFLAGS: -L${SRCDIR}/lib -lmymath2
//#include "mymath.h"
import "C"
import "fmt"

func TestCallDynamicMethod(a, b int) int {
	//#cgo LDFLAGS: -L${SRCDIR}/lib -lmymath2
	res, _ := C.add(C.int(a), C.int(b))
	fmt.Println("Test call dynamic method: 所有用法同静态库调用")
	return int(res)
}
