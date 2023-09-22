package cgo_function_variables

/*
#cgo CFLAGS: -I../include
#include "common.h"
extern void testFunctionVariable(uintptr_t h, char* str);
*/
import "C"
import "runtime/cgo"

//export testFunctionVariable
func testFunctionVariable(h C.uintptr_t, str *C.char) {
	f := cgo.Handle(h).Value().(func(*C.char))
	f(str)
}

func CallbackFunction(str *C.char) {
	C.puts(C.CString("the secong step in go side: "))
	C.puts(str)
}
