package cgo_function_variables

/*
#cgo CFLAGS: -I../include
#include "common.h"
extern char* callbackFunc(char* str);
*/
import "C"
import "fmt"

//export callbackFunc
func callbackFunc(str *C.char) *C.char {
	fmt.Printf("the input params : %s\n", C.GoString(str))
	return C.CString("returned value")
}
