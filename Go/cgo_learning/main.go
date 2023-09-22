package main

/*
#cgo CFLAGS: -I./include
#include "func.h"
*/
import "C"
import (
	_ "cgo_learning/cgo_function_variables"
	"cgo_learning/cgo_type_conversion"
)

func main() {
	//cgo_call_cmethod.TestCallCMethod("hello world!")

	//addres, _ := cgo_call_lib.TestCallStaticLib(4, 5)
	//addres, _ := TestCallStaticLib(4, 5)
	//fmt.Printf("调用静态库add()结果为:%v\n", addres)

	//addres := cgo_call_lib.TestCallDynamicMethod(4, 5)
	//fmt.Printf("调用动态库add()结果为:%v\n", addres)

	//use cgo to wrap the c function qsort()
	//slice := []int64{42, 9, 101, 95, 27, 25}
	//go_sort.SortByqsort(slice, func(a, b int) bool {
	//	return slice[a] > slice[b]
	//})
	//fmt.Println(slice)

	//use export instruction to export go function to c function
	//cgo_export_gomethod.TestExportGoMethodAsCMethod()

	//type conversion
	//cgo_type_conversion.TestTypeConversion()

	//pass the go slice pointer to c
	cgo_type_conversion.TestPassingGoPointer()

	//function uintptr_t callback
	//handle := cgo.NewHandle(cgo_function_variables.CallbackFunction)
	//C.workflow(C.uintptr_t(handle), C.CString("evething you input"))
	//handle.Delete()

	//function pointer callback
	//C.workflowForFuncPointer(C.testfunctionPointer(unsafe.Pointer(C.gatewayFunc)), C.CString("Hello world!"))

}
