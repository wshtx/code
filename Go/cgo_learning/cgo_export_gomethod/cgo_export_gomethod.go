package cgo_export_gomethod

/*
#include "test.h"
extern void TestExportGoMethodAsCMethod();
*/
import "C"
import "fmt"

//export TestExportGoMethodAsCMethod
func TestExportGoMethodAsCMethod() {
	fmt.Println("Test export go method as c method: 使用export指令导出函数时，该导出函数所在文件的注释中不能有其他函数的实现/全局变量的声明, 可以有函数声明/头文件")
}
