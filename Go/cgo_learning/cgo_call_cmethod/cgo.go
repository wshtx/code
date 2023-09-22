package cgo_call_cmethod

import "C"

//#include "test.h"
import "C"
import (
	"fmt"
)

func TestCallCMethod(str string) {
	C.TestCMethod(C.CString("Test C Method"))
	C.TestGoRewriteCMethod(C.CString("Test Go rewrite the C Method"))
}

//export TestGoRewriteCMethod
func TestGoRewriteCMethod(s *C.char) {
	fmt.Println(C.GoString(s))
}
